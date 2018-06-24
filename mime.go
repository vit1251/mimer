package mimer

import "io"
import "fmt"
import "mime/multipart"

/** writeHeaders writes the Mime-Version, Date, Reply-To, From, To and Subject headers.
 */
func (m *Mail) writeHeaders(buf io.Writer) error {

	if _, err := buf.Write([]byte(m.fromHeader())); err != nil {
		return err
	}

	if _, err := buf.Write([]byte("Mime-Version: 1.0\r\n")); err != nil {
		return err
	}

	fmt.Fprintf(buf, "Date: %s\r\n", m.date)

	if m.replyTo != "" {
		fmt.Fprintf(buf, "Reply-To: %s\r\n", m.replyTo)
	}

	fmt.Fprintf(buf, "Subject: %s\r\n", m.subject)

	for _, to := range m.toAddrs {
		fmt.Fprintf(buf, "To: %s\r\n", to)
	}

	for _, cc := range m.ccAddrs {
		fmt.Fprintf(buf, "CC: %s\r\n", cc)
	}

	if m.writeBccHeader {
		for _, bcc := range m.bccAddrs {
			fmt.Fprintf(buf, "BCC: %s\r\n", bcc)
		}
	}

	return nil
}


/** fromHeader returns a correctly formatted From header, optionally with a name
 * component.
 */
func (m *Mail) fromHeader() string {

	if m.fromName == "" {
		return fmt.Sprintf("From: %s\r\n", m.fromAddr)
	}

	return fmt.Sprintf("From: %s <%s>\r\n", m.fromName, m.fromAddr)
}

/** writeBody writes the text/plain and text/html mime parts.
 */
func (m *Mail) writeBody(w io.Writer, boundary string) error {

	alt := multipart.NewWriter(w)
	defer alt.Close()

}
