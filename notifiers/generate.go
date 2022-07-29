//go:build ignore
// +build ignore

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/statping-ng/statping-ng/utils"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/html"
)

var (
	mjmlApplication string
	mjmlPrivate     string
)

func main() {
	utils.InitEnvs()

	mjmlApplication = os.Getenv("MJML_APP")
	mjmlPrivate = os.Getenv("MJML_PRIVATE")

	if mjmlApplication == "" || mjmlPrivate == "" {
		fmt.Println("skipping email MJML template render, missing MJML_APP and MJML_PRIVATE")
		return
	}

	fmt.Println("Generating success/failure email templates from MJML to a HTML golang constant")

	success := convertMJML(emailSuccessMJML)
	fail := convertMJML(emailFailureMJML)

	htmlOut := `// DO NOT EDIT ** This file was generated with go generate on ` + utils.Now().String() + ` ** DO NOT EDIT //
package notifiers

const emailSuccess = ` + minimize(success) + `

const emailFailure = ` + minimize(fail) + `

`

	utils.SaveFile("email_rendered.go", []byte(htmlOut))

	fmt.Println("Email MJML to HTML const saved: notifiers/email_rendered.go")
}

type mjmlInput struct {
	Mjml string `json:"mjml"`
}

func minimize(val string) string {
	m := minify.New()
	m.Add("text/html", &html.Minifier{
		KeepDefaultAttrVals: true,
	})
	s, err := m.String("text/html", val)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("`%s`", s)
}

func convertMJML(mjml string) string {
	input, _ := json.Marshal(mjmlInput{mjml})
	auth := fmt.Sprintf("%s:%s", mjmlApplication, mjmlPrivate)
	resp, _, err := utils.HttpRequest("https://"+auth+"@api.mjml.io/v1/render", "POST", "application/json", nil, bytes.NewBuffer(input), 15*time.Minute, false, nil)
	if err != nil {
		panic(err)
	}
	var respData mjmlApi
	if err := json.Unmarshal(resp, &respData); err != nil {
		panic(err)
	}
	return respData.Html
}

type mjmlApi struct {
	Html    string `json:"html"`
	Mjml    string `json:"mjml"`
	Version string `json:"mjml_version"`
}

const emailFailureMJML = `<mjml>
  <mj-head>
    <mj-title>Statping Service Notification</mj-title>
  </mj-head>
  <mj-body background-color="#E7E7E7">
    <mj-raw>
      <!-- Top Bar -->
    </mj-raw>
    <mj-section background-color="#a30911" background-url="https://assets.statping.com/offlinebanner.png" padding="0px">
      <mj-column>
        <mj-image width="45px" href="https://statping.com" src="https://assets.statping.com/iconlight.png" align="center" alt="Sphero"></mj-image>
      </mj-column>
    </mj-section>
  
    <mj-section background-color="#ffffff">
      <mj-column width="100%">
        <mj-text font-family="Ubuntu, Helvetica, Arial, sans-serif" font-size="22px" padding="15px" line-height="30px">
         {{.Service.Name}} is currently offline, you might want to check it.
        </mj-text>
        
        
            <mj-section padding-left="0px" padding-right="0px" padding-top="10px" padding-bottom="10px">
        <mj-column>
          <mj-text font-color="#d50d0d" align="center" font-size="20px" color="#626262">Offline for {{.Service.Downtime.Human}}</mj-text>
          
          <mj-button border-radius="4px" background-color="#cb121c" href="{{.Core.Domain}}/service/{{.Service.Id}}">View Dashboard</mj-button>
          
    </mj-column>
  </mj-section>
        
       
    <mj-raw>
      <!-- Bottom Graphic -->
    </mj-raw>
    
        
        
   
  <mj-section padding="0px" background-color="#fafafa">
        <mj-column>
          <mj-text font-size="20px" color="#626262">Service Domain</mj-text>
          <mj-text padding-top="0px" font-size="14px" color="#626262">{{.Service.Domain}}</mj-text>
    </mj-column>
  </mj-section>
    
    <mj-section padding="0px"  background-color="#ffffff">
        <mj-column>
          <mj-text font-size="20px" color="#626262">Current Issue</mj-text>
          <mj-text padding-top="0px" font-size="14px" color="#626262">{{.Failure.Issue}}</mj-text>
    </mj-column>
  </mj-section>
    
   
 <mj-spacer height="30px" />
        
    <mj-section padding="0" background-url="https://assets.statping.com/offlinebanner.png" background-color="#a30911">
      <mj-column>
        <mj-image width="250px" href="https://statping.com" src="https://assets.statping.com/statpingcom.png" align="center"></mj-image>
      </mj-column>
    </mj-section>
        
         <mj-section padding-bottom="0" padding-top="10px">
        <mj-column>
        <mj-text color="#445566" font-size="11px" align="center" line-height="16px">
            You are receiving this email because one of your services has changed on your Statping instance. You can modify this email on the Email Notifier page in Settings.
          </mj-text>
          <mj-text color="#445566" font-size="11px" align="center" line-height="16px">
            &copy; Statping
          </mj-text>
        </mj-column>
      </mj-section>
        
         <mj-section padding-top="0" padding-bottom="0">
        <mj-group>
          <mj-column width="100%" padding-right="0">
            <mj-text color="#445566" font-size="11px" align="center" line-height="16px" font-weight="bold">
              <a class="footer-link" href="https://statping.com">Statping.com</a>&#xA0;&#xA0;&#xA0;&#xA0;&#xA0;&#xA0;&#xA0;&#xA0; 
              
              <a class="footer-link" href="https://github.com/statping/statping">Github</a>&#xA0;&#xA0;&#xA0;&#xA0;&#xA0;&#xA0;&#xA0;&#xA0;
              
              <a class="footer-link" href="https://statping.com/privacy">Privacy</a>&#xA0;&#xA0;&#xA0;&#xA0;&#xA0;&#xA0;&#xA0;&#xA0;
            </mj-text>
          </mj-column>
        </mj-group>

      </mj-section>
      </mj-column>
    </mj-section>
        
  </mj-body>
</mjml>`

const emailSuccessMJML = `<mjml>
  <mj-head>
    <mj-title>Statping Service Notification</mj-title>
  </mj-head>
  <mj-body background-color="#E7E7E7">
    <mj-raw>
      <!-- Top Bar -->
    </mj-raw>
    <mj-section background-color="#12ab0c" background-url="https://assets.statping.com/greenbackground.png" padding="0px">
      <mj-column>
        <mj-image width="45px" href="https://statping.com" src="https://assets.statping.com/iconlight.png" align="center" alt="Sphero"></mj-image>
      </mj-column>
    </mj-section>
  
    <mj-section background-color="#ffffff">
      <mj-column width="100%">
        <mj-text font-family="Ubuntu, Helvetica, Arial, sans-serif" font-size="22px" padding="15px" line-height="30px">
         {{.Service.Name}} is currently offline, you might want to check it.
        </mj-text>
        
        
            <mj-section padding-left="0px" padding-right="0px" padding-top="10px" padding-bottom="10px">
        <mj-column>
          <mj-text font-color="#d50d0d" align="center" font-size="20px" color="#626262">Offline for {{.Service.Downtime.Human}}</mj-text>
          
          <mj-button border-radius="4px" background-color="#4caf50" href="{{.Core.Domain}}/service/{{.Service.Id}}">View Dashboard</mj-button>
          
    </mj-column>
  </mj-section>
        
       
    <mj-raw>
      <!-- Bottom Graphic -->
    </mj-raw>
    
        
        
   
  <mj-section padding="0px" background-color="#fafafa">
        <mj-column>
          <mj-text font-size="20px" color="#626262">Service Domain</mj-text>
          <mj-text padding-top="0px" font-size="14px" color="#626262">{{.Service.Domain}}</mj-text>
    </mj-column>
  </mj-section>
    
    <mj-section padding="0px"  background-color="#ffffff">
        <mj-column>
          <mj-text font-size="20px" color="#626262">Current Issue</mj-text>
          <mj-text padding-top="0px" font-size="14px" color="#626262">{{.Failure.Issue}}</mj-text>
    </mj-column>
  </mj-section>
    
   
 <mj-spacer height="30px" />
        
    <mj-section padding="0" background-url="https://assets.statping.com/greenbackground.png" background-color="#12ab0c">
      <mj-column>
        <mj-image width="250px" href="https://statping.com" src="https://assets.statping.com/statpingcom.png" align="center"></mj-image>
      </mj-column>
    </mj-section>
        
         <mj-section padding-bottom="0" padding-top="10px">
        <mj-column>
        <mj-text color="#445566" font-size="11px" align="center" line-height="16px">
            You are receiving this email because one of your services has changed on your Statping instance. You can modify this email on the Email Notifier page in Settings.
          </mj-text>
          <mj-text color="#445566" font-size="11px" align="center" line-height="16px">
            &copy; Statping
          </mj-text>
        </mj-column>
      </mj-section>
        
         <mj-section padding-top="0" padding-bottom="0">
        <mj-group>
          <mj-column width="100%" padding-right="0">
            <mj-text color="#445566" font-size="11px" align="center" line-height="16px" font-weight="bold">
              <a class="footer-link" href="https://statping.com">Statping.com</a>&#xA0;&#xA0;&#xA0;&#xA0;&#xA0;&#xA0;&#xA0;&#xA0; 
              
              <a class="footer-link" href="https://github.com/statping/statping">Github</a>&#xA0;&#xA0;&#xA0;&#xA0;&#xA0;&#xA0;&#xA0;&#xA0;
              
              <a class="footer-link" href="https://statping.com/privacy">Privacy</a>&#xA0;&#xA0;&#xA0;&#xA0;&#xA0;&#xA0;&#xA0;&#xA0;
            </mj-text>
          </mj-column>
        </mj-group>

      </mj-section>
      </mj-column>
    </mj-section>
        
  </mj-body>
</mjml>`
