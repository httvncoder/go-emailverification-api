package controllers

/*
 * Todo: Single function
 */

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/smtp"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/julienschmidt/httprouter"
	"github.com/sebastianbroekhoven/go-emailverification-api/models"
)

type (
	// EmailVerifyController type struc
	EmailVerifyController struct{}
)

// NewEmailVerifyController controller
func NewEmailVerifyController() *EmailVerifyController {
	return &EmailVerifyController{}
}

// EmailVerify function
func (evc EmailVerifyController) EmailVerify(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	h := new(models.Message)
	h.Question.JobTime = time.Now()

	emailaddress := p.ByName("emailaddress")

	if govalidator.IsEmail(emailaddress) == false {
		log.Println("Failed: Not a valid e-mail address.")
		h.Question.JobEmailAddress = emailaddress
		h.Question.JobStatus = "Failed"
		h.Question.JobMessage = "Not a valid e-mail address"
		// Marshal provided interface into JSON structure
		hj, _ := json.MarshalIndent(h, "", "    ")
		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", hj)
		return
	}

	h.Answer.EmailAddress = emailaddress
	emailuser, emaildomain := ParseEmailAddress(emailaddress)
	h.Answer.EmailUser = emailuser
	h.Answer.EmailDomain = emaildomain

	mx, err := net.LookupMX(emaildomain)
	if err != nil {
		log.Println("Failed: No MX records found.")
		log.Println(err)
		h.Question.JobEmailAddress = emailaddress
		h.Question.JobStatus = "Failed"
		h.Question.JobMessage = "Failed: No MX records found."
		// Marshal provided interface into JSON structure
		hj, _ := json.MarshalIndent(h, "", "    ")
		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", hj)
		return
	}

	h.Answer.MXRecords = mx

	c, err := smtp.Dial(mx[0].Host + ":25")
	if err != nil {
		log.Println("Failed: Cannot connect to MX: " + mx[0].Host)
		log.Println(err)
		h.Question.JobEmailAddress = emailaddress
		h.Question.JobStatus = "Failed"
		h.Question.JobMessage = "Failed: Cannot connect to MX: " + mx[0].Host
		// Marshal provided interface into JSON structure
		hj, _ := json.MarshalIndent(h, "", "    ")
		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", hj)
		return
	}

	err = c.Hello("blaas.io")
	if err != nil {
		log.Println("Failed: No hello from: " + mx[0].Host)
		log.Println(err)
		h.Question.JobEmailAddress = emailaddress
		h.Question.JobStatus = "Failed"
		h.Question.JobMessage = "Failed: No hello from: " + mx[0].Host
		// Marshal provided interface into JSON structure
		hj, _ := json.MarshalIndent(h, "", "    ")
		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", hj)
		return
	}

	err = c.Mail("verify-email-address@blaas.io")
	if err != nil {
		log.Println("Failed: No mail command: " + mx[0].Host)
		log.Println(err)
		h.Question.JobEmailAddress = emailaddress
		h.Question.JobStatus = "Failed"
		h.Question.JobMessage = "Failed: No mail command: " + mx[0].Host
		// Marshal provided interface into JSON structure
		hj, _ := json.MarshalIndent(h, "", "    ")
		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", hj)
		return
	}

	if err := c.Rcpt(emailaddress); err != nil {
		log.Println("OK: Job done!")
		log.Println(err)
		h.Question.JobEmailAddress = emailaddress
		h.Question.JobStatus = "OK"
		h.Question.JobMessage = "OK: Job done!"
		h.Answer.ValidationResult = "invalid"
		// Marshal provided interface into JSON structure
		hj, _ := json.MarshalIndent(h, "", "    ")
		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", hj)
		err = c.Quit()
		return
	}

	log.Println("OK: Job done!")
	h.Question.JobEmailAddress = emailaddress
	h.Question.JobStatus = "OK"
	h.Question.JobMessage = "OK: Job done!"
	h.Answer.ValidationResult = "valid"
	// Marshal provided interface into JSON structure
	hj, _ := json.MarshalIndent(h, "", "    ")
	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(200)
	fmt.Fprintf(w, "%s", hj)
}

// ParseEmailAddress for splitting
func ParseEmailAddress(emailaddress string) (string, string) {
	split := strings.Split(emailaddress, "@")
	if len(split) == 2 {
		return split[0], split[1]
	}
	return "", ""
}
