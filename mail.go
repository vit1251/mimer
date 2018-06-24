package mimer

import "regexp"
import "time"
import "fmt"
import "mime"

type Mail struct {
	html  BodyPart
	plain BodyPart

	toAddrs        []string
	ccAddrs        []string
	bccAddrs       []string
	subject        string
	fromAddr       string
	fromName       string
	replyTo        string
	attachments    []Attachment
	trimRegex      *regexp.Regexp
	writeBccHeader bool
	date string
}

/** New returns an instance of Mail.
 *
 *	mail := mimer.New()
 *
 */
func New() *Mail {
	mail := &Mail{
		trimRegex:      regexp.MustCompile("\r?\n"),
		writeBccHeader: false,
		date:           time.Now().Format(time.RFC1123Z),
	}
	return mail
}

/** To sets a list of recipient addresses.
 *
 * You can pass one or more addresses to this method, all of which are viewable to the recipients.
 *
 *	mail.To("dom@itsallbroken.com", "another@itsallbroken.com")
 *
 * or pass a slice of strings:
 *
 *	tos := []string{
 *		"one@itsallbroken.com",
 *		"two@itsallbroken.com"
 *	}
 *	mail.To(tos...)
 */
func (m *Mail) To(addrs ...string) {
	m.toAddrs = []string{}
	for _, addr := range addrs {
		trimmed := m.trimRegex.ReplaceAllString(addr, "")
		if trimmed == "" {
			continue
		}
		m.toAddrs = append(m.toAddrs, trimmed)
	}
}

/** Bcc sets a list of blind carbon copy (BCC) addresses.
 *
 * You can pass one or more addresses to this method, none of which are viewable to the recipients.
 *
 *	mail.Bcc("dom@itsallbroken.com", "another@itsallbroken.com")
 *
 * or pass a slice of strings:
 *
 *	bccs := []string{
 *		"one@itsallbroken.com",
 *		"two@itsallbroken.com"
 *	}
 *
 *	mail.Bcc(bccs...)
 */
func (m *Mail) Bcc(addrs ...string) {
	m.bccAddrs = []string{}
	for _, addr := range addrs {
		trimmed := m.trimRegex.ReplaceAllString(addr, "")
		if trimmed == "" {
			continue
		}
	m.bccAddrs = append(m.bccAddrs, trimmed)
	}
}

/** WriteBccHeader writes the BCC header to the MIME body when true. Defaults to false.
 *
 * This is usually required when writing the MIME body to an email API such as
 * ( Amazon's SES, but can cause problems when sending emails via a SMTP server.
 *
 * Specifically, RFC822 says:
 *
 *	Some  systems  may choose to include the text of the "Bcc" field only in the
 *	author(s)'s  copy,  while  others  may also include it in the text sent to
 *	all those indicated in the "Bcc" list.
 *
 * This ambiguity can result in some SMTP servers not stripping the BCC header
 * and exposing the BCC addressees to recipients.
 */
func (m *Mail) WriteBccHeader(shouldWrite bool) {
    m.writeBccHeader = shouldWrite
}

/** Cc sets a list of carbon copy (CC) addresses.
 *
 * You can pass one or more addresses to this method, which are viewable to the other recipients.
 *
 *	mail.Cc("dom@itsallbroken.com", "another@itsallbroken.com")
 *
 * or pass a slice of strings:
 *
 *	ccs := []string{
 *		"one@itsallbroken.com",
 *		"two@itsallbroken.com"
 *	}
 *	mail.Cc(ccs...)
 */
func (m *Mail) Cc(addrs ...string) {
	m.ccAddrs = []string{}
	for _, addr := range addrs {
		trimmed := m.trimRegex.ReplaceAllString(addr, "")
		if trimmed == "" {
			continue
		}
		m.ccAddrs = append(m.ccAddrs, trimmed)
	}
}

/** From sets the sender email address.
 *
 * Users should also consider setting FromName().
 */
func (m *Mail) From(addr string) {
	m.fromAddr = m.trimRegex.ReplaceAllString(addr, "")
}

/** FromName sets the sender name.
 *
 * If set, emails typically display as being from:
 *
 *	"From Name" <sender@example.com>
 *
 * If name contains non-ASCII characters, it is Q-encoded according to RFC1342.
 */
func (m *Mail) FromName(name string) {
	m.fromName = mime.QEncoding.Encode("UTF-8", m.trimRegex.ReplaceAllString(name, ""))
}

/** ReplyTo sets the Reply-To email address.
 *
 * Setting a ReplyTo address is optional.
 */
func (m *Mail) ReplyTo(addr string) {
	m.replyTo = m.trimRegex.ReplaceAllString(addr, "")
}

/** Subject sets the email subject line.
 *
 * If sub contains non-ASCII characters, it is Q-encoded according to RFC1342.
 */
func (m *Mail) Subject(sub string) {
	m.subject = mime.QEncoding.Encode("UTF-8", m.trimRegex.ReplaceAllString(sub, ""))
}

/** String returns a redacted description of the mail state, typically for logging or debugging purposes.
 */
func (m *Mail) String() string {
	return fmt.Sprintf("&Mail{date: %q}", m.date)
}
