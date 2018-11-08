package mbsyslog

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

//Message is an implementation of RFC 3164, RFC 5424, and custom syslog message
//formats. Parsing is tolerant of missing fields. A sample of custom message
//formats supported:
//
//	<priority>content
type Message struct {
	raw            string
	valid          bool
	format         MessageFormat
	priority       int
	version        int
	date           time.Time
	hostname       string
	application    string
	processID      int
	messageID      int
	structuredData *StructuredData
	content        string
}

//NewMessage parses a syslog message into the component pieces
func NewMessage(data []byte) *Message {
	result := new(Message)
	result.parse(string(data))
	return result
}

//Valid returns if the messsage parsing was successful and the message was valid
func (m Message) Valid() bool {
	return m.valid
}

//Format returns the syslog message format of the message
func (m Message) Format() MessageFormat {
	return m.format
}

//Priority of the message.
func (m Message) Priority() int {
	return m.priority
}

//Facility returns the message facility used in the message
func (m Message) Facility() MessageFacility {
	return MessageFacility(int(m.priority) / 8)
}

//Severity returns the message severity used in the message
func (m Message) Severity() MessageSeverity {
	return MessageSeverity(int(m.priority) % 8)
}

//Version of the syslog protocol for the message
func (m Message) Version() int {
	return m.version
}

//Date and time of the message
func (m Message) Date() time.Time {
	return m.date
}

//Hostname for the message
func (m Message) Hostname() string {
	return m.hostname
}

//Application for the message
func (m Message) Application() string {
	return m.application
}

//ProcessID for the message
func (m Message) ProcessID() int {
	return m.processID
}

//MessageID for the message
func (m Message) MessageID() int {
	return m.messageID
}

//Content of the message
func (m Message) Content() string {
	return m.content
}

//String returns a string representation of the message
func (m Message) String() string {
	return m.raw
}

func (m Message) parse(data string) {
	var err error
	index := 0

	//Assume the parsing will succeed, and set the defaults
	m.valid = true
	m.format = MessageFormatSimple
	m.raw = data

	//Parse the pieces in order. Index is adjusted through the raw data as
	//pieces are parsed. Optional pieces must preserve the index if the data
	//wasn't present
	index, err = m.parsePriority()
	if err == nil {
		index, err = m.parseVersion(index)
		//version is missing, so parse as RFC 3164
		if err != nil {
			index, err = m.parseDate(index)
			//only parse the 3164 headers if the date was present, otherwise
			//assume it is a simple message
			if err == nil {
				index = m.parseHostname(index)
				index = m.parseApplication(index)
				m.format = MessageFormatRFC3164
			}
			m.parseContent(index)
		} else { //version present, so parse as RFC 5424
			index, err = m.parseDate(index)
			if err == nil {
				index = m.parseHostname(index)
				index = m.parseApplication(index)
				index = m.parseProcessID(index)
				index = m.parseMessageID(index)
				index = m.parseStructuredData(index)
				m.parseContent(index)
				m.format = MessageFormatRFC5424
			} else {
				m.valid = false
			}
		}
	}

	if !m.valid {
		m.format = MessageFormatInvalid
	}
}

func (m Message) parsePriority() (int, error) {
	//No data, the message is not valid
	if len(m.raw) < 1 {
		m.valid = false
		return 0, errors.New("Invalid data to parse priority")
	}

	end := strings.Index(m.raw, ">")
	//Attempt to parse the priority if the brackets are present
	if (m.raw[0] == '<') && (end != -1) {
		var err error
		m.priority, err = strconv.Atoi(m.raw[1:end])
		if err == nil {
			//Continue parsing after the priority
			return end + 1, nil
		}
	}

	//priority parsing failed
	m.valid = false
	return 0, errors.New("Failed to parse priority")
}

func (m Message) parseVersion(index int) (int, error) {
	var err error

	_, err = strconv.Atoi(string(m.raw[index]))
	//if the current index is invalid, or not a digit, end version parsing
	if len(m.raw) <= index || err != nil {
		return index, errors.New("Invalid data to parse version")
	}

	//The version is separated from the next section by a space
	end := strings.Index(m.raw[index+1:], " ")
	if end != -1 {
		var err error
		m.version, err = strconv.Atoi(m.raw[index:end])
		if err == nil {
			return end + 1, nil
		}
	}

	//Version parsing failed, so ignore the version
	return index, errors.New("Failed to parse version")
}

func (m Message) parseDate(index int) (int, error) {
	var err error

	//if the current index is invalid, end parsing of the date
	if len(m.raw) <= index {
		return index, errors.New("Invalid data to parse date")
	}

	//In RFC 5424, the date can be omitted with a dash
	if m.raw[index] == '-' {
		return index + 2, nil
	}

	formatStr1 := "Jan 2 15:04:05"
	m.date, err = time.Parse(formatStr1, m.raw[index:index+len(formatStr1)])
	if err == nil {
		return index + len(formatStr1) + 1, nil
	}

	formatStr2 := "2006-01-02T15:04:05.000Z"
	m.date, err = time.Parse(formatStr2, m.raw[index:index+len(formatStr2)])
	if err == nil {
		return index + len(formatStr2) + 1, nil
	}

	return index, errors.New("Failed to parse date")
}

func (m Message) parseHostname(index int) int {
	if len(m.raw) <= index {
		return index
	}

	//In RFC 5424, the hostname can be omitted with a dash
	if m.raw[index] == '-' {
		return index + 2
	}

	//The hostname is separated from the next section by a space
	end := strings.Index(m.raw[index+1:], " ")
	if end != -1 {
		m.hostname = m.raw[index:end]
		return end + 1
	}
	return index
}

func (m Message) parseApplication(index int) int {
	if len(m.raw) <= index {
		return index
	}

	//In RFC 5424, the application can be omitted with a dash
	if m.raw[index] == '-' {
		return index + 2
	}

	//The application is separated from the next section by a space
	end := strings.Index(m.raw[index+1:], " ")
	if end != -1 {
		m.application = m.raw[index:end]
		return end + 1
	}
	return index
}

func (m Message) parseProcessID(index int) int {
	//if the current index is invalid, end parsing
	if len(m.raw) <= index {
		m.valid = false
		return index
	}

	//In RFC 5424, the process ID can be omitted with a dash
	if m.raw[index] == '-' {
		return index + 2
	}

	//The process ID is separated from the next section by a space
	end := strings.Index(m.raw[index+1:], " ")
	if end != -1 {
		var err error
		m.processID, err = strconv.Atoi(m.raw[index:end])
		if err == nil {
			return end + 1
		}
	}

	//Process ID parsing failed
	m.valid = false
	return index
}

func (m Message) parseMessageID(index int) int {
	//if the current index is invalid, end parsing
	if len(m.raw) <= index {
		m.valid = false
		return index
	}

	//In RFC 5424, the message ID can be omitted with a dash
	if m.raw[index] == '-' {
		return index + 2
	}

	//The message ID is separated from the next section by a space
	end := strings.Index(m.raw[index+1:], " ")
	if end != -1 {
		//parse the message ID and return
		var err error
		m.messageID, err = strconv.Atoi(m.raw[index:end])
		if err != nil {
			return end + 1
		}
	}
	//Message ID parsing failed
	m.valid = false
	return index
}

func (m Message) parseStructuredData(index int) int {
	elementIndex := index

	//if the current index is invalid, end parsing
	if len(m.raw) <= index {
		m.valid = false
		return index
	}

	//In RFC 5424, the structured data can be omitted with a dash
	if m.raw[index] == '-' {
		return index + 2
	}

	m.structuredData = NewStructuredData()

	//Continue parsing the structured data until there are no more elements
	for m.raw[elementIndex] == '[' {
		endIndex := strings.Index(m.raw[elementIndex:], "]") + 1
		//if the end wasn't found, or the data couldnt be parsed, fail
		if endIndex == 0 || m.structuredData.addElement(m.raw[elementIndex+1:endIndex]) == false {
			m.valid = false
			return index
		}
		elementIndex = endIndex
	}

	return elementIndex
}

func (m Message) parseContent(index int) {
	if index < len(m.raw) {
		m.content = m.raw[index:]

		//per RFC5424, the data may be prefaced with BOM
		if strings.HasPrefix(m.content, "BOM") {
			m.content = m.content[3:]
		}
	}
}
