package mimer

import "io"

type attachment struct {
	filename string
	content  io.Reader
	inline   bool
}

/** Attach adds the contents of r to the email as an attachment with name as the filename.
 *
 * r is not read until Send is called.
 */
func (m *Mail) Attach(name string, r io.Reader) {
	m.attachments = append(m.attachments, attachment{
		filename: name,
		content:  r,
		inline:   false,
	})
}

/** AttachInline adds the contents of r to the email as an inline attachment.
 * Inline attachments are typically used within the email body, such as a logo
 * or header image. It is up to the user to ensure name is unique.
 *
 * Files can be referenced by their name within the email using the cid URL
 * protocol:
 *
 *	<img src="cid:myFileName"/>
 *
 * r is not read until Send is called.
 */
func (m *Mail) AttachInline(name string, r io.Reader) {
	m.attachments = append(m.attachments, attachment{
		filename: name,
		content:  r,
		inline:   true,
	})
}

/** ClearAttachments removes all current attachments.
 */
func (m *Mail) ClearAttachments() {
	m.attachments = []attachment{}
}
