<template>
<div class="col-12">
<div class="row mb-4">
	
	<div class="col">
	<h4 class="h4 mb-2">Overview</h4>
		
			<a @click.prevent='tab="Features"' class="d-block mb-1 text-link" href="#">Features</a>
		
			<a @click.prevent='tab="Paid Hosting"' class="d-block mb-1 text-link" href="#">Paid Hosting</a>
		
			<a @click.prevent='tab="Types of Monitoring"' class="d-block mb-1 text-link" href="#">Types of Monitoring</a>
		
			<a @click.prevent='tab="Notifiers"' class="d-block mb-1 text-link" href="#">Notifiers</a>
		
			<a @click.prevent='tab="Bulk Import Services"' class="d-block mb-1 text-link" href="#">Bulk Import Services</a>
		
			<a @click.prevent='tab="Environment Variables"' class="d-block mb-1 text-link" href="#">Environment Variables</a>
		
			<a @click.prevent='tab="Testing"' class="d-block mb-1 text-link" href="#">Testing</a>
		
			<a @click.prevent='tab="Badges"' class="d-block mb-1 text-link" href="#">Badges</a>
		
			<a @click.prevent='tab="Contributing"' class="d-block mb-1 text-link" href="#">Contributing</a>
		
			<a @click.prevent='tab="Mobile App"' class="d-block mb-1 text-link" href="#">Mobile App</a>
		
	</div>
	
	<div class="col">
	<h4 class="h4 mb-2">Installation</h4>
		
			<a @click.prevent='tab="Docker"' class="d-block mb-1 text-link" href="#">Docker</a>
		
			<a @click.prevent='tab="Docker-Compose"' class="d-block mb-1 text-link" href="#">Docker-Compose</a>
		
			<a @click.prevent='tab="AWS EC2"' class="d-block mb-1 text-link" href="#">AWS EC2</a>
		
			<a @click.prevent='tab="Mac"' class="d-block mb-1 text-link" href="#">Mac</a>
		
			<a @click.prevent='tab="Linux"' class="d-block mb-1 text-link" href="#">Linux</a>
		
			<a @click.prevent='tab="Windows"' class="d-block mb-1 text-link" href="#">Windows</a>
		
			<a @click.prevent='tab="Cloud Foundry"' class="d-block mb-1 text-link" href="#">Cloud Foundry</a>
		
	</div>
	
	<div class="col">
	<h4 class="h4 mb-2">Extras</h4>
		
			<a @click.prevent='tab="Development"' class="d-block mb-1 text-link" href="#">Development</a>
		
			<a @click.prevent='tab="Makefile"' class="d-block mb-1 text-link" href="#">Makefile</a>
		
			<a @click.prevent='tab="GraphQL"' class="d-block mb-1 text-link" href="#">GraphQL</a>
		
			<a @click.prevent='tab="SSL"' class="d-block mb-1 text-link" href="#">SSL</a>
		
			<a @click.prevent='tab="Config with .env File"' class="d-block mb-1 text-link" href="#">Config with .env File</a>
		
			<a @click.prevent='tab="config.yml"' class="d-block mb-1 text-link" href="#">config.yml</a>
		
			<a @click.prevent='tab="Static HTML Exporter"' class="d-block mb-1 text-link" href="#">Static HTML Exporter</a>
		
			<a @click.prevent='tab="Prometheus Exporter"' class="d-block mb-1 text-link" href="#">Prometheus Exporter</a>
		
			<a @click.prevent='tab="PGP Signature"' class="d-block mb-1 text-link" href="#">PGP Signature</a>
		
	</div>
	
</div>

<div class="col-12" v-if='tab === "Home"'>
	<div v-pre>
		<p align="center">
<img width="60%" src="https://s3-us-west-2.amazonaws.com/gitimgs/statping.png">
<img width="70%" src="https://img.cjx.io/statuptokenbalance.gif">
</p>

<p align="center">
<img alt="GitHub commit activity" src="https://img.shields.io/github/commit-activity/m/statping/statping"> 
<a href="https://github.com/statping/statping/labels/priority"><img src="https://img.shields.io/github/issues/statping/statping/priority?color=blue&label=priority"></a> <a href="https://github.com/statping/statping/labels/bug"><img src="https://img.shields.io/github/issues/statping/statping/bug?color=red&label=bugs"></a> <a href="https://github.com/statping/statping/labels/feature"><img src="https://img.shields.io/github/issues/statping/statping/feature?color=pink&label=features"></a>
</p>

<h1>Docker</h1>

<p>View the <a href="https://github.com/statping/statping/wiki/Docker" target="_blank">Docker Wiki</a> for more information and advanced usage.</p>

<pre><code class="language-sh">docker run -it -p 8080:8080 statping/statping
</code></pre>

<h1>MacOSX</h1>

<p>View the <a href="https://github.com/statping/statping/wiki/Mac" target="_blank">Mac Wiki</a> for more information and other methods.</p>

<pre><code class="language-sh">brew tap statping/statping
brew install statping
</code></pre>

<h1>Linux</h1>

<p>View the <a href="https://github.com/statping/statping/wiki/Linux" target="_blank">Linux Wiki</a> for more information and other methods.</p>

<pre><code class="language-sh">curl -o- -L https://statping.com/install.sh | bash
</code></pre>

<p><a href="https://snapcraft.io/statping" target="_blank"><img src="https://snapcraft.io/static/images/badges/en/snap-store-white.svg" alt="Get it from the Snap Store" /></a></p>

<pre><code class="language-sh">snap install statping
</code></pre>

<h1>Windows</h1>

<p>Download the <a href="https://github.com/statping/statping/releases/latest" target="_blank">latest version</a> as an <code>.exe</code> from <a href="https://github.com/statping/statping/releases/latest" target="_blank">Releases</a>. View the <a href="https://github.com/statping/statping/wiki/Windows" target="_blank">Windows Wiki</a> for more information and other methods.</p>

<h1>EC2</h1>

<p>// exploring new methods for the AMI image.</p>

<h1>Development</h1>

	</div>
</div>

	<div class="col-12" v-if='tab === "Features"'>
		<h1 class="h1 mt-5 mb-5 text-muted">Features</h1>
		<span class="spacer"></span>
		<div v-pre>
				<h1>3 Different Databases</h1>

<p>You can use MySQL, Postgres, or SQLite as a database for your Statping status page. The server will automatically upgrade your database tables depending on which database you have.</p>

<h1>Easy to Startup</h1>

<p>Statping is an extremely easy to setup website monitoring tool without fussing with dependencies or packages. Simply download and install the precompile binary for your operating system. Statping works on Windows, Mac, Linux, Docker, and even the Raspberry Pi.</p>

<h1>Plugins</h1>

<p>Statping is an awesome Status Page generator that allows you to create your own plugins with Golang Plugins! You don&rsquo;t need to request a PR or even tell us about your plugin. Plugin&rsquo;s are compiled and then send as a binary to the Statping <code>/plugins</code> folder. Test your plugins using the <code>statup test plugin</code> command, checkout the <a href="https://github.com/statping/statping/wiki/Statping-Plugins" target="_blank">Plugin Wiki</a> to see detailed information about creating plugins.</p>

<h1>No Maintenance</h1>

<p>Many other website monitoring applications will collect data until the server fails because of hard drive is 100% full. Statping will automatically delete records to make sure your server will stay UP for years. The EC2 AMI Image is a great way to host your status page without worrying about it crashing one day. Statping will automatically upgrade its software when you reboot your computer.</p>

<h1>Email &amp; Slack Notifications</h1>

<p>Receive email notifications if your website or application goes offline. Statping includes SMTP connections so you can use AWS SES, or any other SMTP emailing service. Go in the Email Settings in Settings to configure these options.</p>

<h1>Prometheus Exporter</h1>

<p>If you want a deeper view of your applications status, you can use Grafana and Prometheus to graph all types of data about your services. Read more about the <a href="https://github.com/statping/statping/wiki/Prometheus-Exporter" target="_blank">Prometheus Exporter</a></p>

		</div>
	</div>

	<div class="col-12" v-if='tab === "Paid Hosting"'>
		<h1 class="h1 mt-5 mb-5 text-muted">Paid Hosting</h1>
		<span class="spacer"></span>
		<div v-pre>
				<p><a href="https://statping.com/" target="_blank">Statping.com</a> now provides a service that will host your Statping instance for you! For $6.00/month you won&rsquo;t need to worry about maintaining your own Statping instance. Billing platform is using Stripe and you can cancel at anytime.</p>

<p><a href="https://statping.com/"><img with="70%" src="https://assets.statping.com/pricingtable.png"></a></p>

<p>Help support the project by donating or have us host your Statping instance.</p>

		</div>
	</div>

	<div class="col-12" v-if='tab === "Types of Monitoring"'>
		<h1 class="h1 mt-5 mb-5 text-muted">Types of Monitoring</h1>
		<span class="spacer"></span>
		<div v-pre>
				<p>You can monitor your application by using a simple HTTP GET to the endpoint to return back a response and status code. Normally you want a 200 status code on an HTTP request. You might want to require a 404 or 500 error as a response code though. With each service you can include a Timeout in seconds to work with your long running services.</p>

<h1>HTTP Endpoints with Custom POST</h1>

<p>For more advanced monitoring you can add a data as a HTTP POST request. This is useful for automatically submitting JSON, or making sure your signup form is working correctly.</p>

<p align="center">
<img width="100%" src="https://img.cjx.io/statup-httpservice.png">
</p>

<p>With a HTTP service, you can POST a JSON string to your endpoint to retrieve any type of response back. You can then use Regex in the Expected Response field to parse a custom response that exactly matches your status requirements.</p>

<h1>TCP/UDP Services</h1>

<p>For other services that don&rsquo;t use HTTP, you can monitor any type of service by using the PORT of the service. If you&rsquo;re Ethereum Blockchain server is running on 8545, you can use TCP to monitor your server. With a TCP service, you can monitor your Docker containers, or remove service running on a custom port. You don&rsquo;t need to include <code>http</code> in the endpoint field, just IP or Hostname.</p>

<p align="center">
<img width="100%" src="https://img.cjx.io/statup-tcpservice.png">
</p>

<h1>ICMP Service</h1>

<p>You can send a <a href="https://en.wikipedia.org/wiki/Internet_Control_Message_Protocol" target="_blank">ICMP</a> (ping) to an endpoint rather than HTTP/TCP/UDP request for a quick response. For this service is fully work, you may need to run Statping as a <code>sudo</code> user.</p>

		</div>
	</div>

	<div class="col-12" v-if='tab === "Notifiers"'>
		<h1 class="h1 mt-5 mb-5 text-muted">Notifiers</h1>
		<span class="spacer"></span>
		<div v-pre>
				<p align="center">
<img width="80%" src="https://s3-us-west-2.amazonaws.com/gitimgs/statupnotifiers.png">
</p>

<p>Statping includes multiple Notifiers to alert you when your services are offline. You can also create your own notifier and send a Push Request to this repo! Creating a custom notifier is pretty easy as long as you follow the requirements. A notifier will automatically be installed into the users Statping database, and form values will save without any hassles. 💃</p>

<p align="center">
<a href="https://github.com/statping/statping/wiki/Notifier-Example">Example Code</a> | <a href="https://github.com/statping/statping/wiki/Notifier-Events">Events</a> | <a href="https://github.com/statping/statping/tree/master/notifiers">View Notifiers</a></p>

<h2>Notifier Requirements</h2>

<ul>
<li>Must have a unique <code>METHOD</code> name</li>
<li>Struct must have <code>*notifier.Notification</code> pointer in it.</li>
<li>Must create and add your notifier variable in <code>init()</code></li>
<li>Should have a form for user to input their variables/keys. <code>Form: []notifier.NotificationForm</code></li>
</ul>

<h2>Notifier Interface (required)</h2>

<p>Statping has the <code>Notifier</code> interface which you&rsquo;ll need to include in your notifier. Statping includes many other events/triggers for your notifier, checkout <a href="https://github.com/statping/statping/wiki/Notifier-Events">Notifier Events</a> to see all of them.</p>

<pre><code class="language-go">// Notifier interface is required to create a new Notifier
type Notifier interface {
        // OnSuccess is triggered when a service is successful
	OnSuccess(*services.Service) error
        // OnFailure is triggered when a service is failing
	OnFailure(*services.Service, *failures.Failure) error
        // OnTest is triggered for testing
	OnTest() (string, error)
}
</code></pre>

<pre><code class="language-go">var example = &amp;Example{&amp;notifier.Notification{
	Method: &quot;example&quot;,                               // unique method name
	Host:   &quot;http://exmaplehost.com&quot;,                // default 'host' field
	Form: []notifier.NotificationForm{{
		Type:        &quot;text&quot;,                     // text, password, number, or email
		Title:       &quot;Host&quot;,                     // The title of value in form
		Placeholder: &quot;Insert your Host here.&quot;,   // Optional placeholder in input
		DbField:     &quot;host&quot;,                     // An accepted DbField value (read below)
	}},
}
</code></pre>

<h2>Notifier Form</h2>

<p>Include a form with your notifier so other users can save API keys, username, passwords, and other values.</p>

<pre><code class="language-go">// NotificationForm contains the HTML fields for each variable/input you want the notifier to accept.
type NotificationForm struct {
	Type        string `json:&quot;type&quot;`        // the html input type (text, password, email)
	Title       string `json:&quot;title&quot;`       // include a title for ease of use
	Placeholder string `json:&quot;placeholder&quot;` // add a placeholder for the input
	DbField     string `json:&quot;field&quot;`       // true variable key for input
	SmallText   string `json:&quot;small_text&quot;`  // insert small text under a html input
	Required    bool   `json:&quot;required&quot;`    // require this input on the html form
	IsHidden    bool   `json:&quot;hidden&quot;`      // hide this form element from end user
	IsList      bool   `json:&quot;list&quot;`        // make this form element a comma separated list
	IsSwitch    bool   `json:&quot;switch&quot;`      // make the notifier a boolean true/false switch
}
</code></pre>

<h3>Example Notifier Form</h3>

<p>This is the Slack Notifier <code>Form</code> fields.</p>

<pre><code class="language-go">Form: []notifier.NotificationForm{{
		Type:        &quot;text&quot;,
		Title:       &quot;Incoming webhooker Url&quot;,
		Placeholder: &quot;Insert your slack webhook URL here.&quot;,
		SmallText:   &quot;Incoming webhooker URL from &lt;a href=\&quot;https://api.slack.com/apps\&quot; target=\&quot;_blank\&quot;&gt;slack Apps&lt;/a&gt;&quot;,
		DbField:     &quot;Host&quot;,
		Required:    true,
	}}
}
</code></pre>

<h3>Accepted DbField Values</h3>

<p>The <code>notifier.NotificationForm</code> has a field called <code>DbField</code> which is the column to save the value into the database. Below are the acceptable DbField string names to include in your form.
- <code>host</code> used for a URL or API endpoint
- <code>username</code> used for a username
- <code>password</code> used for a password
- <code>port</code> used for a integer port number
- <code>api_key</code> used for some kind of API key
- <code>api_secret</code> used for some API secret
- <code>var1</code> used for any type of string
- <code>var2</code> used for any type of string (extra)</p>

<h3>Form Elements</h3>

<p>You can completely custom your notifications to include a detailed form.
- <code>Type</code> is a HTML input type for your field
- <code>Title</code> give your input element a title
- <code>Placeholder</code> optional field if you want a placeholder in input
- <code>DbField</code> required field to save variable into database (read above)
- <code>Placeholder</code> optional field for inserting small hint under the input</p>

<h1>Adding Notifiers</h1>

<p>To add a notifier to the Statping application, simply append your Notifier in the <code>AttachNotifiers()</code> function inside of <a href="https://github.com/statping/statping/blob/master/core/core.go" target="_blank">core/core.go</a>.</p>

<pre><code class="language-go">// AttachNotifiers will attach all the notifier's into the system
func AttachNotifiers() error {
	return notifier.AddNotifiers(
		notifiers.Command,
		notifiers.Discorder,
		notifiers.Emailer,
		notifiers.LineNotify,
		notifiers.Mobile,
		notifiers.Slacker,
		notifiers.Telegram,
		notifiers.Twilio,
		notifiers.Webhook,
	)
}
</code></pre>

<h6><a href="https://github.com/statping/statping/blob/master/core/core.go#L183" target="_blank">AttachNotifiers</a></h6>

		</div>
	</div>

	<div class="col-12" v-if='tab === "Bulk Import Services"'>
		<h1 class="h1 mt-5 mb-5 text-muted">Bulk Import Services</h1>
		<span class="spacer"></span>
		<div v-pre>
				<p>You can import multiple services when Statping first loads by creating a file named <a href="https://github.com/statping/statping/wiki/services.yml" target="_blank">services.yml</a> in the working directory for Statping. It will insert the new service into the database, and will not be re-added on reboot. All services must be an array under the <code>services:</code> field.</p>

<h2>Dynamic Yaml</h2>

<p>With Yaml, you can insert &ldquo;anchors&rdquo; to make receptive fields simple! Checkout the example below. The <code>&amp;tcpservice</code> anchor will return all the fields belonging to <code>x-tcpservice:</code>. To reuse these fields for each service, you can insert <code>&lt;&lt;: *tcpservice</code> and thats it!</p>

<pre><code class="language-yaml">x-tcpservice: &amp;tcpservice
  type: tcp
  check_interval: 60
  timeout: 15
  allow_notifications: true
  notify_after: 0
  notify_all_changes: true
  public: true
  redirect: true

x-httpservice: &amp;httpservice
  type: http
  method: GET
  check_interval: 45
  timeout: 10
  expected_status: 200
  allow_notifications: true
  notify_after: 2
  notify_all_changes: true
  public: true
  redirect: true

services:

  - name: Statping Demo
    domain: https://demo.statping.com
    &lt;&lt;: *httpservice

  - name: Portainer
    domain: portainer
    port: 9000
    &lt;&lt;: *tcpservice

  - name: Statping Github
    domain: https://github.com/statping/statping
    &lt;&lt;: *httpservice
</code></pre>

<h2>Duplicates</h2>

<p>Services that have already been added into the database will not be added again. This is done by creating a SHA256 hash of the following fields&hellip;</p>

<pre><code>sha256(name:EXAMPLEdomain:HTTP://DOMAIN.COMport:8080type:HTTPmethod:GET)
</code></pre>

		</div>
	</div>

	<div class="col-12" v-if='tab === "Environment Variables"'>
		<h1 class="h1 mt-5 mb-5 text-muted">Environment Variables</h1>
		<span class="spacer"></span>
		<div v-pre>
				<p>This page will list all environment variables that could be passed into Statping. You can see all default values in <a href="https://github.com/statping/statping/blob/dev/utils/env.go" target="_blank">utils/env.go</a> file under the <code>InitEnvs()</code> function.</p>

<h3>Statping Server</h3>

<ul>
<li><code>PORT</code>                      - Set the outgoing port for the HTTP server (or use &ndash;port, default: 8080)</li>
<li><code>HOST</code>                        - Bind a specific IP address to the HTTP server (or use &ndash;ip)</li>
<li><code>VERBOSE</code>                   - Display more logs in verbose mode. (1 - 4)</li>
<li><code>STATPING_DIR</code>              - Set a absolute path for the root path of Statping server (logs, assets, SQL db)</li>
<li><code>DISABLE_LOGS</code>              - Disable viewing and writing to the log file (default is false)</li>
<li><code>GO_ENV</code>                    - Run Statping in testmode, will bypass HTTP authentication (if set as &lsquo;test&rsquo;)</li>
<li><code>REMOVE_AFTER</code>              - Automatically delete records after time (default 3 months, &lsquo;12h = 12 hours&rsquo;)</li>
<li><code>CLEANUP_INTERVAL</code>          - Interval to check for old records (default 1 hour, &lsquo;1h = 1 hour&rsquo;)</li>
<li><code>ALLOW_REPORTS</code>             - Send Statping anonymous <a href="https://sentry.io/" target="_blank">error reports</a> so we can see issues (default is false)</li>
<li><code>SERVER_PORT</code>               - Port number to run Statping HTTP server on (or use -p/&ndash;port)</li>
</ul>

<h3>Automatic SSL Certificate</h3>

<p>With LetsEncrypt enabled, Statping will run through the SSL process and create the SSL certs in the <code>certs</code> folder. Read more about the SSL Process on the <a href="https://github.com/statping/statping/wiki/SSL" target="_blank">SSL Wiki</a>.
- <code>LETSENCRYPT_ENABLE</code>        - Set to <strong>true</strong> to have LetsEncrypt enabled. (defaults to false)
- <code>LETSENCRYPT_HOST</code>          - Domain to generate SSL certificate
- <code>LETSENCRYPT_EMAIL</code>         - Email address that gets sent with the LetsEncrypt Request
- <code>LETSENCRYPT_LOCAL</code>         - Set for LetsEncrypt testing</p>

<h3>Database</h3>

<ul>
<li><code>DB_CONN</code>                   - Database connection (sqlite, postgres, mysql) <strong>Will automatically start if set to &lsquo;sqlite&rsquo;</strong></li>
<li><code>DB_HOST</code>                   - Database hostname or IP address</li>
<li><code>DB_USER</code>                   - Database username</li>
<li><code>DB_PASS</code>                   - Database password</li>
<li><code>DB_PORT</code>                   - Database port (5432, 3306, &hellip;)</li>
<li><code>DB_DATABASE</code>               - Database connection&rsquo;s database name</li>
<li><code>POSTGRES_SSLMODE</code>          - Enable <a href="https://www.postgresql.org/docs/9.1/libpq-ssl.html" target="_blank">Postgres SSL Mode</a> &lsquo;ssl_mode=VALUE&rsquo; (enable/disable/verify-full/verify-ca)</li>
<li><code>MAX_OPEN_CONN</code>             - Set <a href="https://golang.org/pkg/database/sql/#DB.SetMaxOpenConns" target="_blank">Maximum Open Connections</a> for database server (default: 25)</li>
<li><code>MAX_IDLE_CONN</code>             - Set <a href="https://golang.org/pkg/database/sql/#DB.SetMaxIdleConns" target="_blank">Maximum Idle Connections</a> for database server (default: 25)</li>
<li><code>MAX_LIFE_CONN</code>             - Set <a href="https://golang.org/pkg/database/sql/#DB.SetConnMaxLifetime" target="_blank">Maximum Life Connections</a> for database server (default: 25)</li>
<li><code>PREFIX</code> - Add a prefix string to each Prometheus metric (default is empty)</li>
</ul>

<h3>Connection</h3>

<ul>
<li><p><code>BASE_PATH</code>                 - Set the base URL prefix (set to &lsquo;monitor&rsquo; if URL is domain.com/monitor)</p></li>

<li><p><code>PREFIX</code>                    - A Prefix for each value in <a href="https://github.com/statping/statping/wiki/Prometheus-Exporter" target="_blank">Prometheus /metric exporter</a></p></li>

<li><p><code>HTTP_PROXY</code>                - Use a <a href="https://golang.org/pkg/net/http/#ProxyFromEnvironment" target="_blank">HTTP Proxy</a> for HTTP Requests</p></li>

<li><p><code>AUTH_USERNAME</code>             - HTTP Basic Authentication username</p></li>

<li><p><code>AUTH_PASSWORD</code>             - HTTP Basic Authentication password</p></li>

<li><p><code>DISABLE_HTTP</code>              - Disable HTTP server if set to true</p></li>

<li><p><code>DISABLE_COLORS</code>            - Disable colors in terminal logs if set to true</p></li>

<li><p><code>DEBUG</code>                     - Enables pprof golang debugging on port 9090</p></li>

<li><p><code>LOGS_MAX_COUNT</code>            - Maximum amount of log files (defaults to 5)</p></li>

<li><p><code>LOGS_MAX_AGE</code>              - Maximum age for log files (defaults to 28 days)</p></li>

<li><p><code>LOGS_MAX_SIZE</code>             - Maximum size for log files (defaults to 16 MB)</p></li>

<li><p><code>LANGUAGE</code>                  - Language to use (en, fr, ru, more to come&hellip;)</p>

<h3>Assets</h3></li>

<li><p><code>SASS</code>                      - Set the absolute path to the <a href="https://sass-lang.com/" target="_blank">sass</a> binary location (find with <code>which sass</code>)</p></li>

<li><p><code>USE_ASSETS</code>                - Automatically use assets from &lsquo;assets folder&rsquo; (true/false)</p></li>
</ul>

<h3>Automatic Fills</h3>

<ul>
<li><code>NAME</code>                      - Set a name for the Statping status page</li>
<li><code>DESCRIPTION</code>               - Set a description for the Statping status page</li>
<li><code>DOMAIN</code>                    - Set a URL for the Statping status page</li>
<li><code>ADMIN_USER</code>                - Username for administrator account (default: admin)</li>
<li><code>ADMIN_PASSWORD</code>            - Password for administrator account (default: admin)</li>
<li><code>API_SECRET</code>                - Set a custom API Secret for API Authentication</li>
<li><code>SAMPLE_DATA</code>               - Insert sample services, groups and more (default: true)</li>
</ul>

		</div>
	</div>

	<div class="col-12" v-if='tab === "Testing"'>
		<h1 class="h1 mt-5 mb-5 text-muted">Testing</h1>
		<span class="spacer"></span>
		<div v-pre>
				<p>If you want to test your updates with the current golang testing units, you can follow the guide below to run a full test process. Each test for Statping will run in MySQL, Postgres, and SQlite to make sure all database types work correctly.</p>

<h2>Requirements</h2>

<ul>
<li>Golang (1.14.* or whichever <code>go.mod</code> states) <img src="https://img.shields.io/github/go-mod/go-version/statping/statping" alt="GitHub go.mod Go version" />

<ul>
<li><a href="https://github.com/GeertJohan/go.rice" target="_blank">go.rice</a> Compiles all assets into golang file</li>
</ul></li>
<li>Node (currently using <code>10.17</code>)

<ul>
<li>Yarn</li>
<li>Sass</li>
</ul></li>
</ul>

<h2>Basic Testing</h2>

<p>The easiest way to run local unit testing by using SQLite and following the commands below&hellip;
1. Make sure you have the requirements above. (being able to run <code>rice</code>)
2. <code>cd frontend &amp;&amp; yarn</code> Install npm packages for the frontend
3. <code>make compile</code> Creates production version of Vue frontend and compiles all assets with <code>rice</code>.
4. <code>make install</code> Will install <code>statping</code> locally so your system test CLI commands/flags.
5. <code>go test -v ./...</code> Run all tests for Statping.</p>

<blockquote>
<p>If you run Statping on port 8585, and then run <code>yarn dev</code> the debug Vue server will be running on port 8888.</p>
</blockquote>

<h2>Notifier Testing</h2>

<p>To test the notifiers, you must include an environment variable that is set in each notifier test. If that notifier test environment variable is not set, it will be skipped.</p>

<h2>Create Docker Databases</h2>

<p>For more advanced testing, you can connect to a different type of database by using the <code>DB_CONN</code> environment variable. You can start them with Docker. Docker is available for Linux, Mac and Windows. You can download/install it by going to the <a href="https://docs.docker.com/install/" target="_blank">Docker Installation</a> site.</p>

<pre><code class="language-bash">docker run -it -d \
   -p 3306:3306 \
   -env MYSQL_ROOT_PASSWORD=password123 \
   -env MYSQL_DATABASE=root mysql
</code></pre>

<pre><code class="language-bash">docker run -it -d \
   -p 5432:5432 \
   -env POSTGRES_PASSWORD=password123 \
   -env POSTGRES_USER=root \
   -env POSTGRES_DB=root postgres
</code></pre>

<p>Once you have MySQL and Postgres running, you can begin the testing. SQLite database will automatically create a <code>statping.db</code> file and will delete after testing.</p>

<h2>Run Tests</h2>

<p>Insert the database environment variables to auto connect the the databases and run the normal test command: <code>go test -v ./...</code>. You&rsquo;ll see a verbose output of each test. If all tests pass, make a push request! 💃</p>

<pre><code class="language-bash">DB_CONN=sqlite \
DB_DATABASE=root \
DB_USER=root \
DB_PASS=password123 \
DB_HOST=localhost \
go test -v ./...
</code></pre>

		</div>
	</div>

	<div class="col-12" v-if='tab === "Badges"'>
		<h1 class="h1 mt-5 mb-5 text-muted">Badges</h1>
		<span class="spacer"></span>
		<div v-pre>
				<p>Using shields.io you can easily create Badges for your services. Using the URL below, you can format your own SVG badge:</p>

<pre><code>https://img.shields.io/badge/dynamic/json?url=&lt;URL&gt;&amp;label=&lt;LABEL&gt;&amp;query=&lt;$.DATA.SUBDATA&gt;&amp;color=&lt;COLOR&gt;&amp;prefix=&lt;PREFIX&gt;&amp;suffix=&lt;SUFFIX&gt;
</code></pre>

<ul>
<li><code>url</code> - URL for JSON endpoint</li>
<li><code>label</code> - Text for the left side of the badge</li>
<li><code>query</code> - <a href="https://jsonpath.com/" target="_blank">JSON Path</a> to be set for right side value</li>
</ul>

<h2>Examples</h2>

<h4>Service Uptime Percent</h4>

<p><img src="https://img.shields.io/badge/dynamic/json?color=blue&label=%20Statping%20Uptime&query=%24.online_24_hours&url=https%3A%2F%2Fdemo.statping.com%2Fapi%2Fservices%2F1&suffix=%"></p>

<ul>
<li>URL: <a href="https://demo.statping.com/api/services/1" target="_blank">https://demo.statping.com/api/services/1</a></li>
<li>JSON Path: <code>$.online_24_hours</code></li>
<li>Suffix: <code>%</code></li>
</ul>

<pre><code>https://img.shields.io/badge/dynamic/json?color=blue&amp;label=%20Statping%20Uptime&amp;query=%24.online_24_hours&amp;url=https%3A%2F%2Fdemo.statping.com%2Fapi%2Fservices%2F1&amp;suffix=%
</code></pre>

<h4>Count Services</h4>

<p><img src="https://img.shields.io/badge/dynamic/json?color=purple&label=Demo%20Site&query=%24.services&url=https://demo.statping.com/health&suffix=%20services"></p>

<ul>
<li>URL: <a href="https://demo.statping.com/health" target="_blank">https://demo.statping.com/health</a></li>
<li>JSON Path: <code>$.services</code></li>
<li>Suffix: <code>services</code></li>
</ul>

<pre><code>https://img.shields.io/badge/dynamic/json?color=purple&amp;label=Demo%20Site&amp;query=%24.services&amp;url=https://demo.statping.com/health&amp;suffix=%20services
</code></pre>

		</div>
	</div>

	<div class="col-12" v-if='tab === "Contributing"'>
		<h1 class="h1 mt-5 mb-5 text-muted">Contributing</h1>
		<span class="spacer"></span>
		<div v-pre>
				<p>Have a feature you want to implement into Statping!? Awesome! Follow this guide
to see how you can test, compile and build Statping for production use. I
recommend you use <code>make</code> with this process, it will save you time and it will
auto-include many customized parameters to get everything working correctly.</p>

<blockquote>
<p>Send all PR&rsquo;s to <code>dev</code> branch!</p>
</blockquote>

<h1>Requirements</h1>

<ul>
<li>Golang <code>1.14</code></li>
<li>Node <code>10.15.3</code></li>
</ul>

<h1>Dependencies</h1>

<p>Statping has a couple of required dependencies when testing and compiling the
binary. The <a href="https://github.com/statping/statping/blob/master/Makefile" target="_blank">Makefile</a>
will make these tasks a lot easier. Take a look at the Makefile to see what
commands are running. Run the command below to get setup right away.</p>

<pre><code class="language-bash">make clean compile build
</code></pre>

<p>List of requirements for compiling assets, building binary, and testing.
- <a href="https://golang.org/" target="_blank">Go Language</a>
- <a href="https://docs.docker.com/" target="_blank">Docker</a>
- <a href="https://sass-lang.com/install" target="_blank">SASS</a>
- <a href="https://www.cypress.io/" target="_blank">Cypress</a> (only used for UI testing, <code>make cypress-install</code>)</p>

<h1>Compiling Assets</h1>

<p>This Golang project uses <a href="https://github.com/GeertJohan/go.rice" target="_blank">rice</a> to
compile static assets into a single file. The file <code>source/rice-box.go</code> is never
committed to the Github repo, it is automatically created on the build. Statping
also requires <code>sass</code> to be installed on your local OS. To compile all the
static assets run the command below:</p>

<pre><code class="language-bash">make compile
</code></pre>

<p>After this is complete, you&rsquo;ll notice the <code>source/rice-box.go</code> file has been
generated. You can now continue to build, and test.</p>

<h1>Testing</h1>

<p>Statping includes multiple ways to Test the application, you can run the <code>make</code>
command, or the normal <code>go test</code> command. To see the full experience of your
updates, you can even run Cypress tests which are in the <code>.dev/test</code> folder.</p>

<p>Statping will run all tests in the <code>cmd</code> folder on MySQL, Postgres, and SQLite
databases. You can run <code>make databases</code> to automatically create MySQL and
Postgres with Docker.</p>

<h6>Golang Testing:</h6>

<pre><code class="language-bash">make test
</code></pre>

<h6>Cypress UI Testing:</h6>

<pre><code class="language-bash">make cypress-test
</code></pre>

<h6>Test Everything:</h6>

<pre><code class="language-bash">make test-all
</code></pre>

<h1>Build</h1>

<p>Statping will build on all operating systems except Windows 32-bit. I
personally use <a href="https://github.com/karalabe/xgo" target="_blank">xgo</a> to cross-compile on
multiple systems using Docker. Follow the commands below to build on your local
system.</p>

<h6>Build for local operating system:</h6>

<pre><code class="language-bash">make compile build
</code></pre>

<h1>Compile for Production</h1>

<p>Once you&rsquo;ve tested and built locally, you can compile Statping for all available
operating systems using the command below. This command will require you to have
<a href="http://docker.io/" target="_blank">Docker</a>.</p>

<pre><code class="language-bash">make build-all
</code></pre>

<h1>What Now</h1>

<p>Everything tested, compiled and worked out!? Awesome! 💃 You can now commit your
changes, and submit a Pull Request with the features/bugs you added or removed.</p>

<p>Please remember to also add the Changelog Entry to describe what you have
changed by using<a href="https://gitlab.com/l0nax/changelog-go" target="_blank">changelog-go</a>.</p>

<p>If your PR is related to an Issue (ie. on a bug fix or when adding a new notifier)
don&rsquo;t forget adding the related GitHub Issue ID like this: <code>[#270](https://github.com/statping/statping/issues/270)</code>
Check out the example below.
Adding <em>one</em> changelog entry should be done in a separate Commit - not only
because this is a good Committing practice, also because it&rsquo;s separated by your
Changes.
When you fix/ add/ &hellip; more than one thing than please describe it a good, short
and precise sentence. Because everyone wants a good and nice-looking <code>CHANGELOG.md</code>
where all changes of a new Version are documented.</p>

<h6>Example for fixing a bug in the UI</h6>

<pre><code class="language-bash">~] changelog new &quot;Fix zooming out on graph doesn't load additional data ([#270](https://github.com/statping/statping/issues/270))&quot;
[0] New Feature          (Added)
[1] Bug Fix              (Fixed)
[2] Feature change       (Changed)
[3] New deprecation      (Deprecated)
[4] Feature removal      (Removed)
[5] Security fix         (Security)
[6] Other                (Other)
&gt;&gt; 1
</code></pre>

		</div>
	</div>

	<div class="col-12" v-if='tab === "Mobile App"'>
		<h1 class="h1 mt-5 mb-5 text-muted">Mobile App</h1>
		<span class="spacer"></span>
		<div v-pre>
				<p>Statping has a free mobile app so you can monitor your websites and applications without the need of a computer.</p>

<p><img src="https://img.shields.io/itunes/v/1445513219.svg" alt="iTunes App Store" /></p>

<p align="center">
<a href="https://play.google.com/store/apps/details?id=com.statping"><img src="https://img.cjx.io/google-play.svg"></a>
<a href="https://itunes.apple.com/us/app/apple-store/id1445513219"><img src="https://img.cjx.io/app-store-badge.svg"></a>
</p>

<p align="center">
<img src="https://img.cjx.io/statping_iphone_bk.png">
</p>

		</div>
	</div>

	<div class="col-12" v-if='tab === "Docker"'>
		<h1 class="h1 mt-5 mb-5 text-muted">Docker</h1>
		<span class="spacer"></span>
		<div v-pre>
				<p>Statping is easily ran on Docker with the light weight Alpine linux image. View on <a href="https://hub.docker.com/r/statping/statping" target="_blank">Docker Hub</a>.</p>

<p align="center">
<img alt="Docker Image Version (tag latest semver)" src="https://img.shields.io/docker/v/statping/statping/latest"> 
<img alt="Docker Image Size (tag)" src="https://img.shields.io/docker/image-size/statping/statping/latest?label=:latest&color=green"> <img alt="Docker Image Size (tag)" src="https://img.shields.io/docker/image-size/statping/statping/base?label=:base"> <img alt="Docker Pulls" src="https://img.shields.io/docker/pulls/statping/statping">
</p>

<h1>Volume Mount</h1>

<ul>
<li><code>/app</code> - Working directory in container. This directory will hold all configs, logs, assets, and sqlite database.</li>
</ul>

<h1>Latest Docker Image</h1>

<p>The <code>latest</code> Docker image uses Alpine Linux to keep it ultra small.</p>

<pre><code class="language-bash">docker run -d \
  -p 8080:8080 \
  --restart always \
  statping/statping
</code></pre>

<h1>Mounting Volume</h1>

<p>You can mount a volume to the <code>/app</code> Statping directory. This folder will contain <code>logs</code>, <code>config.yml</code>, and static assets if you want to edit the SCSS/CSS.</p>

<pre><code class="language-bash">docker run -d \
  -p 8080:8080 \
  -v /mydir/statping:/app \
  --restart always \
  statping/statping
</code></pre>

<h1>Attach a SSL Certificate</h1>

<p>When you mount <code>server.crt</code> and <code>server.key</code> to the <code>/app</code> directory, Statping will run a HTTPS server on port 443. Checkout the <a href="https://github.com/statping/statping/wiki/SSL" target="_blank">SSL Wiki</a> documentation to see more information about this.</p>

<pre><code class="language-bash">docker run -d \
  -p 443:443 \
  -v /mydir/domain.crt:/app/server.crt \
  -v /mydir/domain.key:/app/server.key \
  -v /mydir:/app \
  --restart always \
  statping/statping
</code></pre>

<h1>Development Docker Image</h1>

<p>If you want to run Statping that was build from the source, use the <code>dev</code> Docker image.</p>

<pre><code class="language-bash">docker run -d -p 8080:8080 statping/statping:dev
</code></pre>

<h1>Cypress Testing Docker Image</h1>

<p>This Docker image will pull the latest version of Statping and test the web interface with <a href="https://www.cypress.io/" target="_blank">Cypress</a>.</p>

<pre><code class="language-bash">docker run -it -p 8080:8080 statping/statping:cypress
</code></pre>

<h4>Or use Docker Compose</h4>

<p>This Docker Compose file inlcudes NGINX, Postgres, and Statping.</p>

<h3>Docker Compose with NGINX and Postgres</h3>

<p>Once you initiate the <code>docker-compose.yml</code> file below go to <a href="http://localhost" target="_blank">http://localhost</a> and you&rsquo;ll be forwarded to the /setup page.
Database Authentication
- database: <code>postgres</code>
- port: <code>5432</code>
- username: <code>statup</code>
- password: <code>password123</code>
- database: <code>statup</code></p>

<pre><code class="language-yaml">version: '2.3'

services:

  nginx:
    container_name: nginx
    image: jwilder/nginx-proxy
    ports:
      - 0.0.0.0:80:80
      - 0.0.0.0:443:443
    networks:
      - internet
    restart: always
    volumes:
      - /var/run/docker.sock:/tmp/docker.sock:ro
      - ./statup/nginx/certs:/etc/nginx/certs:ro
      - ./statup/nginx/vhost:/etc/nginx/vhost.d
      - ./statup/nginx/html:/usr/share/nginx/html:ro
      - ./statup/nginx/dhparam:/etc/nginx/dhparam
    environment:
      DEFAULT_HOST: localhost

  statup:
    container_name: statup
    image: statping/statping:latest
    restart: always
    networks:
      - internet
      - database
    depends_on:
      - postgres
    volumes:
      - ./statup/app:/app
    environment:
      VIRTUAL_HOST: localhost
      VIRTUAL_PORT: 8080
      DB_CONN: postgres
      DB_HOST: postgres
      DB_USER: statup
      DB_PASS: password123
      DB_DATABASE: statup
      NAME: EC2 Example
      DESCRIPTION: This is a Statping Docker Compose instance

  postgres:
    container_name: postgres
    image: postgres:10
    restart: always
    networks:
      - database
    volumes:
      - ./statup/postgres:/var/lib/postgresql/data
    environment:
      POSTGRES_PASSWORD: password123
      POSTGRES_USER: statup
      POSTGRES_DB: statup

networks:
  internet:
    driver: bridge
  database:
    driver: bridge
</code></pre>

<p>Or a simple wget&hellip;</p>

<pre><code class="language-bash">wget https://raw.githubusercontent.com/statping/statping/master/servers/docker-compose.yml
docker-compose up -d
</code></pre>

<h3>Docker Compose with Automatic SSL</h3>

<p>You can automatically start a Statping server with automatic SSL encryption using this docker-compose file. First point your domain&rsquo;s DNS to the Statping server, and then run this docker-compose command with DOMAIN and EMAIL. Email is for letsencrypt services.</p>

<pre><code class="language-bash">wget https://raw.githubusercontent.com/statping/statping/master/servers/docker-compose-ssl.yml

LETSENCRYPT_HOST=mydomain.com \
    LETSENCRYPT_EMAIL=info@mydomain.com \
    docker-compose -f docker-compose-ssl.yml up -d
</code></pre>

<h4>Full docker-compose with Automatic SSL</h4>

<pre><code class="language-yaml">version: '2.3'

services:

  nginx:
    container_name: nginx
    image: jwilder/nginx-proxy
    ports:
      - 0.0.0.0:80:80
      - 0.0.0.0:443:443
    labels:
      - &quot;com.github.jrcs.letsencrypt_nginx_proxy_companion.nginx_proxy&quot;
    networks:
      - internet
    restart: always
    volumes:
      - /var/run/docker.sock:/tmp/docker.sock:ro
      - ./statup/nginx/certs:/etc/nginx/certs:ro
      - ./statup/nginx/vhost:/etc/nginx/vhost.d
      - ./statup/nginx/html:/usr/share/nginx/html:ro
      - ./statup/nginx/dhparam:/etc/nginx/dhparam
    environment:
      DEFAULT_HOST: ${LETSENCRYPT_HOST}

  letsencrypt:
    container_name: letsencrypt
    image: jrcs/letsencrypt-nginx-proxy-companion
    networks:
      - internet
    restart: always
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock:ro
      - ./statup/nginx/certs:/etc/nginx/certs
      - ./statup/nginx/vhost:/etc/nginx/vhost.d
      - ./statup/nginx/html:/usr/share/nginx/html
      - ./statup/nginx/dhparam:/etc/nginx/dhparam

  statup:
    container_name: statup
    image: statping/statping:latest
    restart: always
    networks:
      - internet
      - database
    depends_on:
      - postgres
    volumes:
      - ./statup/app:/app
    environment:
      VIRTUAL_HOST: ${LETSENCRYPT_HOST}
      VIRTUAL_PORT: 8080
      LETSENCRYPT_HOST: ${LETSENCRYPT_HOST}
      LETSENCRYPT_EMAIL: ${LETSENCRYPT_EMAIL}
      DB_CONN: postgres
      DB_HOST: postgres
      DB_USER: statup
      DB_PASS: password123
      DB_DATABASE: statup
      NAME: SSL Example
      DESCRIPTION: This Status Status Page should be running ${LETSENCRYPT_HOST} with SSL.

  postgres:
    container_name: postgres
    image: postgres:10
    restart: always
    networks:
      - database
    volumes:
      - ./statup/postgres:/var/lib/postgresql/data
    environment:
      POSTGRES_PASSWORD: password123
      POSTGRES_USER: statup
      POSTGRES_DB: statup

networks:
  internet:
    driver: bridge
  database:
    driver: bridge
</code></pre>

		</div>
	</div>

	<div class="col-12" v-if='tab === "Docker-Compose"'>
		<h1 class="h1 mt-5 mb-5 text-muted">Docker-Compose</h1>
		<span class="spacer"></span>
		<div v-pre>
				<p>This page includes multiple docker-compose.yml setups for you to run. There are many environment variables for additional settings and features, checkout the <a href="https://github.com/statping/statping/wiki/Environment-Variables" target="_blank">Environment Variables Wiki</a> to see them all.</p>

<p align="center">
<img alt="Docker Image Version (tag latest semver)" src="https://img.shields.io/docker/v/statping/statping/latest"> 
<img alt="Docker Image Size (tag)" src="https://img.shields.io/docker/image-size/statping/statping/latest?label=:latest&color=green"> <img alt="Docker Image Size (tag)" src="https://img.shields.io/docker/image-size/statping/statping/base?label=:base"> <img alt="Docker Pulls" src="https://img.shields.io/docker/pulls/statping/statping">
</p>

<h3>Database Configurations</h3>

<ul>
<li><p><a href="#basic-sqlite-connection">SQLite</a></p></li>

<li><p><a href="#mysql-connection">MySQL</a></p></li>

<li><p><a href="#postgres">Postgres</a></p>

<h3>Web Serving</h3></li>

<li><p><a href="#traefik">Traefik</a></p></li>

<li><p><a href="#nginx-proxy-pass">Nginx as proxy pass</a></p></li>

<li><p><a href="#nginx-proxy-pass-with-ssl-letencrypt">Nginx with LetEncrypt SSL</a></p>

<h3>Metrics</h3></li>

<li>
<h2><a href="#grafana-dashboard-with-prometheus">Grafana with Prometheus</a></h2></li>
</ul>

<h1>Docker Volumes</h1>

<ul>
<li><code>/app</code> - Main Volume that you should mount to a directory on the host machine</li>
<li><code>/app/config.yml</code> - Config file</li>
<li><code>/app/logs</code> - Logs folder (rotates every 3 days)</li>
<li><code>/app/assets</code> - Folder containing all CSS, SCSS, JS and other static files (if you use custom styling)</li>
<li><code>/app/statping.db</code> - SQLite database file (if you&rsquo;re using SQLite)
&mdash;</li>
</ul>

<h1>Basic SQLite Connection</h1>

<p>This docker-compose will start Statping on port 8080 and automatically run on a SQLite database.</p>

<pre><code class="language-yaml">statping:
  container_name: statping
  image: statping/statping
  restart: always
  ports:
    - 8080:8080
  volumes:
    - ./statping:/app
  environment:
    DB_CONN: sqlite
</code></pre>

<h1>MySQL Connection</h1>

<pre><code class="language-yaml">statping:
  container_name: statping
  image: statping/statping
  restart: always
  depends_on:
    - mysql
  ports:
    - 8080:8080
  volumes:
    - statping_data:/app
  links:
    - mysql
  environment:
    DB_CONN: mysql
    DB_HOST: mysql
    DB_PORT: 3306
    DB_DATABASE: statping
    DB_USER: root
    DB_PASS: password123

mysql:
  image: mysql:5.7
  volumes:
    - mysql_data:/var/lib/mysql
  restart: always
  ports:
    - 3306:3306
  environment:
    MYSQL_ROOT_PASSWORD: password123
    MYSQL_DATABASE: statping
    MYSQL_USER: root
    MYSQL_PASSWORD: password
</code></pre>

<h1>Postgres</h1>

<pre><code class="language-yaml">statping:
  container_name: statping
  image: statping/statping
  restart: always
  ports:
    - 8080:8080
  volumes:
    - statping_data:/app
  links:
    - postgres
  depends_on:
    - postgres
  environment:
    DB_CONN: postgres
    DB_HOST: postgres
    DB_PORT: 5432
    DB_DATABASE: statping
    DB_USER: root
    DB_PASS: password123

postgres:
  container_name: postgres
  image: postgres:10.0-alpine
  ports:
    - 5432:5432
  volumes:
    - pg_data:/var/lib/postgresql/data/pg_data
  environment:
    POSTGRES_PASSWORD: password123
    POSTGRES_DB: statping
    POSTGRES_USER: root
    POSTGRES_PORT: 5432
    PGDATA: /var/lib/postgresql/data/pg_data
</code></pre>

<h1>Nginx Proxy Pass</h1>

<p>The <a href="https://github.com/jwilder/nginx-proxy" target="_blank">jwilder/nginx-proxy</a> Docker image will automatically redirect a domain to a specific container.</p>

<pre><code class="language-yaml">statping:
  container_name: statping
  image: statping/statping
  restart: always
  volumes:
    - statping_data:/app
  environment:
    DB_CONN: sqlite
    VIRTUAL_HOST: demo.statping.com
    VIRTUAL_PORT: 8080

nginx:
  container_name: nginx
  image: jwilder/nginx-proxy
  ports:
    - 80:80
  volumes:
    - /var/run/docker.sock:/tmp/docker.sock:ro
</code></pre>

<h1>Nginx Proxy Pass with SSL (jwilder/nginx-proxy)</h1>

<p>Using the <a href="https://github.com/jwilder/nginx-proxy" target="_blank">jwilder/nginx-proxy</a> image along with <a href="https://github.com/JrCs/docker-letsencrypt-nginx-proxy-companion" target="_blank">jrcs/docker-letsencrypt-nginx-proxy-companion</a> you&rsquo;ll be able to automatically generate SSL certificates for your Statping instance.</p>

<pre><code class="language-yaml">statping:
  container_name: statping
  image: statping/statping
  restart: always
  volumes:
    - statping_data:/app
  environment:
    DB_CONN: sqlite
    LETSENCRYPT_HOST: demo.statping.com
    VIRTUAL_HOST: demo.statping.com
    VIRTUAL_PORT: 8080

nginx:
  container_name: nginx
  image: jwilder/nginx-proxy
  restart: always
  ports:
    - 80:80
    - 443:443
  volumes:
    - /var/run/docker.sock:/tmp/docker.sock:ro
    - cert_volume:/etc/nginx/certs
    - vhost_volume:/etc/nginx/vhost.d
    - html_volume:/usr/share/nginx/html
  environment:
    DEFAULT_HOST: demo.statping.com
  labels:
    - &quot;com.github.jrcs.letsencrypt_nginx_proxy_companion.nginx_proxy=true&quot;

nginx-ssl:
  container_name: nginx-ssl
  image: jrcs/letsencrypt-nginx-proxy-companion
  restart: always
  links:
    - nginx
  volumes:
    - /var/run/docker.sock:/var/run/docker.sock:ro
    - cert_volume:/etc/nginx/certs
    - vhost_volume:/etc/nginx/vhost.d
    - html_volume:/usr/share/nginx/html
  environment:
    DEFAULT_EMAIL: info@mydomain.com
    NGINX_PROXY_CONTAINER: nginx
</code></pre>

<h1>Grafana Dashboard with Prometheus</h1>

<p>Grafana is an awesome metric visualizer that allows you to create some awesome dashboards. We&rsquo;ve already created a <a href="https://grafana.com/grafana/dashboards/6950" target="_blank">Grafana Dashboard</a> that you can easy import! Checkout the <a href="https://github.com/statping/statping/wiki/Prometheus-Exporter" target="_blank">Grafana Wiki</a> and the <a href="https://github.com/statping/statping/wiki/Prometheus-Exporter" target="_blank">Prometheus Exporter Wiki</a> for more details.</p>

<h5><code>prometheus.yml</code> config file</h5>

<p>This file should be mounted to <code>/etc/prometheus/prometheus.yml</code> in the Prometheus container.</p>

<pre><code class="language-yaml">global:
  scrape_interval:     30s
  evaluation_interval: 30s

scrape_configs:
  - job_name: 'statping'
    scrape_interval: 30s
    bearer_token: 'SECRET API KEY HERE'
    static_configs:
      - targets: ['statping:8080']
</code></pre>

<blockquote>
<p>Be sure to replace <code>SECRET API KEY HERE</code> with your API Secret on the Settings page.</p>
</blockquote>

<h5><code>docker-compose.yml</code> file for Statping, Grafana and Prometheus</h5>

<pre><code class="language-yaml">statping:
  container_name: statping
  image: statping/statping
  restart: always
  volumes:
    - statping_data:/app
  environment:
    DB_CONN: sqlite
    LETSENCRYPT_HOST: demo.statping.com
    VIRTUAL_HOST: demo.statping.com
    VIRTUAL_PORT: 8080

prometheus:
  container_name: prometheus
  image: prom/prometheus:v2.0.0
  restart: always
  ports:
    - 9090:9090
  volumes:
    - prometheus.yml:/etc/prometheus/prometheus.yml # config file is above ^
    - prometheus_data:/prometheus
  links:
    - statping
  depends_on:
    - statping

grafana:
  container_name: grafana
  image: grafana/grafana
  restart: always
  ports:
    - 3000:3000
  volumes:
    - grafana_data:/var/lib/grafana
  environment:
    - GF_SECURITY_ADMIN_PASSWORD=password123
    - GF_USERS_ALLOW_SIGN_UP=false
  depends_on:
    - prometheus
  links:
    - prometheus
</code></pre>

<h1>Traefik</h1>

<p>This docker-compose stack will start Statping and <a href="https://docs.traefik.io" target="_blank">Traefik</a> v2.2. You may need to research the Traefik label configs and edit to your own needs. The Traefik server will run on port 80 and 443, and Statping on port 8080.</p>

<pre><code class="language-yaml">version: '3.4'

services:

  traefik:
    container_name: traefik
    image: traefik:v2.2.0
    restart: always
    command:
      - --entrypoints.web.address=:80
      - --entrypoints.websecure.address=:443
      - --entrypoints.web.http.redirections.entryPoint.to=websecure
      - --entrypoints.web.http.redirections.entryPoint.scheme=https
      - --entrypoints.websecure.http.tls.certResolver=leresolverresolver
      - --providers.docker
      - --log=true
      - --log.level=INFO
    ports:
      - &quot;80:80&quot;
      - &quot;443:443&quot;
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock:ro
    networks:
      - frontend
    labels:
      - &quot;traefik.enable=true&quot;
      - &quot;traefik.http.routers.traefik.rule=Host(`traefik.statping.com`)&quot;
      - &quot;traefik.http.routers.traefik.service=api@internal&quot;
      - &quot;traefik.http.routers.traefik.middlewares=admin&quot;
      - &quot;traefik.http.routers.traefik.tls.certresolver=leresolver&quot;
      - &quot;traefik.http.routers.traefik.entrypoints=websecure&quot;
      - &quot;traefik.http.routers.http-catchall.entrypoints=web&quot;
      - &quot;traefik.http.routers.http-catchall.rule=HostRegexp(`{host:.+}`)&quot;
      - &quot;traefik.http.routers.http-catchall.middlewares=redirect-to-https&quot;
      - &quot;traefik.http.middlewares.redirect-to-https.redirectscheme.scheme=https&quot;

  statping:
    container_name: statping
    image: statping/statping:latest
    restart: always
    networks:
      - frontend
    volumes:
      - ./statping_data:/app
    environment:
      DB_CONN: sqlite
    labels:
      - &quot;traefik.enable=true&quot;
      - &quot;traefik.http.routers.statping.rule=Host(`demo.statping.com`)&quot;
      - &quot;traefik.http.routers.statping.tls.certresolver=leresolver&quot;
      - &quot;traefik.http.routers.statping.entrypoints=websecure&quot;
      - &quot;traefik.http.services.statping.loadbalancer.server.port=8080&quot;

networks:
  frontend:
    driver: bridge
</code></pre>

		</div>
	</div>

	<div class="col-12" v-if='tab === "AWS EC2"'>
		<h1 class="h1 mt-5 mb-5 text-muted">AWS EC2</h1>
		<span class="spacer"></span>
		<div v-pre>
				<p>Running Statping on the smallest EC2 server is very quick using the AWS AMI Image. The AWS AMI Image will automatically start a Statping Docker container that will automatically update to the latest version. Once the EC2 is booted, you can go to the Public DNS domain to view the Statping installation page. The Statping root folder is located at: <code>/statping</code> on the server.</p>

<h1>AMI Image</h1>

<p>Choose the correct AMI Image ID based on your AWS region.
- us-east-1 <code>ami-09ccd23d9c7afba61</code> (Virginia)
- us-east-2 <code>ami-0c6c9b714a501cdb3</code> (Ohio)
- us-west-1 <code>ami-02159cc1fc701a77e</code> (California)
- us-west-2 <code>ami-007c6990949f5ccee</code> (Oregon)
- eu-central-1 <code>ami-06e252d6d8b0c2f1f</code> (Frankfurt)</p>

<h1>Upgrading Statping</h1>

<p>You can upgrade the Statping executable by running the commands below on your EC2 instance.</p>

<pre><code>VERSION=$(curl -s &quot;https://github.com/statping/statping/releases/latest&quot; | grep -o 'tag/[v.0-9]*' | awk -F/ '{print $2}')
wget https://github.com/statping/statping/releases/download/$VERSION/statping-linux-x64.tar.gz
tar -xvzf statping-linux-x64.tar.gz
chmod +x statping
mv statping /usr/local/bin/statping
</code></pre>

<p>You can test the version number by running <code>statping version</code>.</p>

<h1>Instructions</h1>

<h3>1. Create an EC2 instance from AMI Image</h3>

<p>Go to the main EC2 dashboard and click &lsquo;Launch Instance&rsquo;. Then type <code>Statping</code> inside the search field for &lsquo;Community AMI&rsquo;. Once you&rsquo;ve found it in your region, click Select!</p>

<p><img src="https://img.cjx.io/statpingaws-ami.png"></p>

<h3>2. Get the Public DNS for EC2 Instance</h3>

<p>Copy the &lsquo;Public DNS&rsquo; URL and paste it into your browser.</p>

<p><img src="https://img.cjx.io/statping-aws-ec2.png"></p>

<h3>3. Setup Statping</h3>

<p>Use SQLite if you don&rsquo;t want to connect to a remote MySQL or Postgres database.</p>

<p><img src="https://img.cjx.io/statping-setup.png"></p>

<h1>EC2 Server Features</h1>

<p>Running your Statping server on a small EC2 instance is perfect for most users. Below you&rsquo;ll find some commands to get up and running in seconds.
- Super cheap on the t2.nano (~$4.60 monthly)
- Small usage, 8gb of hard drive
- Automatic SSL certificate if you require it
- Automatic reboot when the server needs it
- Automatic database cleanup, so you&rsquo;ll never be at 100% full.
- Automatic docker containers/images removal</p>

<h2>Create Security Groups</h2>

<p>Using the AWS CLI you can copy and paste the commands below to auto create everything for you. The server opens port 80 and 443.</p>

<pre><code class="language-bash">aws ec2 create-security-group --group-name StatpingPublicHTTP --description &quot;Statping HTTP Server on port 80 and 443&quot;
# will response back a Group ID. Copy ID and use it for --group-id below.
</code></pre>

<pre><code class="language-bash">GROUPS=sg-7e8b830f
aws ec2 authorize-security-group-ingress --group-id $GROUPS --protocol tcp --port 80 --cidr 0.0.0.0/0
aws ec2 authorize-security-group-ingress --group-id $GROUPS --protocol tcp --port 443 --cidr 0.0.0.0/0
</code></pre>

<h2>Create EC2 without SSL</h2>

<p>Once your server has started, go to the EC2 Public DNS endpoint. You should be redirected to /setup to continue your installation process! The database information is already inputed for you.</p>

<pre><code class="language-bash">GROUPS=sg-7e8b830f
KEY=MYKEYHERE
AMI_IMAGE=ami-7be8a103

aws ec2 run-instances \
    --image-id $AMI_IMAGE \
    --count 1 --instance-type t2.nano \
    --key-name $KEY \
    --security-group-ids $GROUPS
</code></pre>

<h3>EC2 Server Specs</h3>

<ul>
<li>t2.nano ($4.60 monthly)</li>
<li>8gb SSD Memory</li>
<li>0.5gb RAM</li>
<li>Docker with Docker Compose installed</li>
<li>Running Statping, NGINX, and Postgres</li>
<li>boot scripts to automatically clean unused containers.</li>
</ul>

		</div>
	</div>

	<div class="col-12" v-if='tab === "Mac"'>
		<h1 class="h1 mt-5 mb-5 text-muted">Mac</h1>
		<span class="spacer"></span>
		<div v-pre>
				<h1>Installing on Mac</h1>

<p>Statping includes an easy to use <a href="https://github.com/statping/homebrew-statping" target="_blank">Homebrew Formula</a> to quick get your Status Page up and running locally. Statping on brew is automatically generated for each new release to master. Install with the commands below,</p>

<h6>Using Homebrew</h6>

<pre><code class="language-bash">brew tap statping/statping
brew install statping
</code></pre>

<h6>Using the Terminal</h6>

<pre><code class="language-shell">curl -o- -L https://statping.com/install.sh | bash
</code></pre>

<p align="center">
<img width="80%" src="https://img.cjx.io/statupbrewinstall.gif">
</p>

<p>Once you&rsquo;ve installed it, checkout which version you have by running <code>statping version</code>.</p>

<h1>Compiling SCSS for Custom Theme</h1>

<p>Statping requires <code>sass</code> to be installed to the local machine to compile SCSS into CSS if you want to use the Custom Theme features.</p>

<ul>
<li>Node: <code>npm install sass -g</code></li>
<li>Ruby: <code>gem install sass</code></li>
</ul>

		</div>
	</div>

	<div class="col-12" v-if='tab === "Linux"'>
		<h1 class="h1 mt-5 mb-5 text-muted">Linux</h1>
		<span class="spacer"></span>
		<div v-pre>
				<h1>Installing on Linux</h1>

<p>Installing and updating Statping on Linux can be done by downloading the latest tar.gz file, unzipping, and running the executable. Or run the install/upgrade script below.</p>

<pre><code class="language-shell">curl -o- -L https://statping.com/install.sh | bash
</code></pre>

<h2>Install using Snapcraft</h2>

<p><a href="https://snapcraft.io/statping" target="_blank"><img src="https://snapcraft.io/static/images/badges/en/snap-store-white.svg" alt="Get it from the Snap Store" /></a></p>

<p>If you are using <a href="https://snapcraft.io/statping" target="_blank">snap</a>, you can simply run this command to install Statping.</p>

<pre><code class="language-shell">sudo snap install statping
</code></pre>

<p><a href="https://snapcraft.io/statping">
<img alt="statping" src="https://snapcraft.io/statping/badge.svg" />
</a></p>

<h1>Compiling SCSS for Custom Theme</h1>

<p>Statping requires <code>sass</code> to be installed to the local machine to compile SCSS into CSS if you want to use the Custom Theme features.</p>

<ul>
<li>Apt: <code>apt install ruby-sass -y</code></li>
<li>Node: <code>npm install sass -g</code></li>
<li>Ruby: <code>gem install sass</code></li>
</ul>

<h2>Systemd Service</h2>

<p>Setting up a systemd service is a great way to make sure your Statping server will automatically reboot when needed. You can use the file below for your service. You should have Statping already installed by this step.</p>

<h6>/etc/systemd/system/statping.service</h6>

<pre><code>[Unit]
Description=Statping Server
After=network.target
After=systemd-user-sessions.service
After=network-online.target

[Service]
Type=simple
Restart=always
ExecStart=/usr/local/bin/statping
WorkingDirectory=/usr/local/bin

[Install]
WantedBy=multi-user.target
</code></pre>

<p>Then you can enable and start your systemd service with:</p>

<pre><code>systemctl daemon-reload

systemctl enable statping.service

systemctl start statping
</code></pre>

<p>You&rsquo;re Statping server will now automatically restart when your server restarts.</p>

<p>Please note that Statping&rsquo;s files (logs, config and database) will be created inside &ldquo;/usr/local/bin/&rdquo; after the service is ran for the first time, unless you change &ldquo;WorkingDirectory&rdquo; configuration option in the systemd service configuration as seen above.</p>

<h2>Raspberry Pi</h2>

<p>You can even run Statping on your Raspberry Pi by installing the precompiled binary from <a href="https://github.com/statping/statping/releases/latest" target="_blank">Latest Releases</a>. For the Raspberry Pi 3 you&rsquo;ll want to download the <code>statping-linux-arm7.tar.gz</code> file. Be sure to change <code>VERSION</code> to the latest version in Releases, and include the &lsquo;v&rsquo;.</p>

<pre><code>VERSION=$(curl -sL &quot;https://github.com/statping/statping/releases/latest&quot; | grep -o 'tag/[v.0-9]*' | awk -F/ '{print $2}' | head -n1)
wget https://github.com/statping/statping/releases/download/$VERSION/statping-linux-arm-7.tar.gz
tar -xvzf statping-linux-arm-7.tar.gz
chmod +x statping
mv statping /usr/local/bin/statping

statping version
</code></pre>

<h2>Alpine Linux</h2>

<p>The Docker image is using the Statping Alpine binary since it&rsquo;s so incredibly small. You can run it on your own alpine image by downloading <code>statping-linux-alpine.tar.gz</code> from <a href="https://github.com/statping/statping/releases/latest" target="_blank">Latest Releases</a>.</p>

		</div>
	</div>

	<div class="col-12" v-if='tab === "Windows"'>
		<h1 class="h1 mt-5 mb-5 text-muted">Windows</h1>
		<span class="spacer"></span>
		<div v-pre>
				<h1>Installing on Windows</h1>

<p>Currently, Statping only works on Windows 64-bit computers. Just download the exe file from <a href="https://github.com/statping/statping/releases/latest" target="_blank">Latest Releases</a> and run it in your command prompt. It will create a HTTP server on port 8080, so you can visit <code>http://localhost:8080</code> to see your Statping Status Page.</p>

<h1>Compiling SCSS for Custom Theme</h1>

<p>Statping requires <code>sass</code> to be installed to the local machine to compile SCSS into CSS if you want to use the Custom Theme features.</p>

<ul>
<li>Node: <code>npm install sass -g</code></li>
<li>Ruby: <code>gem install sass</code></li>
</ul>

<h1>Running Statping as a Service</h1>

<p>To ensure Statping is always running, it can be installed to run as a service on a Windows machine. The easiest way to do that is by using NSSM, the <a href="https://nssm.cc/download" target="_blank">Non-Sucking Service Manager</a>. Download and unzip the compressed file to a location on your machine running Statping to get started:
1. Open an administrative command prompt.
2. Change to the folder that contains the 64 bit version of NSSM.
3. Type &ldquo;nssm install Statping&rdquo; and press enter.
4. For the properties, use the following as an example:
    Path: C:\Program Files\Statping\statping.exe
    Startup directory: C:\Program Files\Statping
5. Click &ldquo;Install&rdquo;.
6. Launch the windows services manager.
7. Run Statping.</p>

<h2>Known Issues with Windows</h2>

<p>Unfortunately, Statping only works on Windows 64-bit processors. If you have more than 4gb of ram, there&rsquo;s a good chance you already have a 64-bit processor. Download the <a href="https://github.com/statping/statping/releases/latest" target="_blank">Latest Releases</a> of Statping, extract the ZIP file, then double click on the <code>statping.exe</code> file. You can use a SQLite database for a quick setup, or connect to a local/remote Postgres or MySQL database server.</p>

		</div>
	</div>

	<div class="col-12" v-if='tab === "Cloud Foundry"'>
		<h1 class="h1 mt-5 mb-5 text-muted">Cloud Foundry</h1>
		<span class="spacer"></span>
		<div v-pre>
				<h3>Installation</h3>

<p>For OSX, you can install the <a href="https://github.com/cloudfoundry/cli" target="_blank">Cloud Foundry CLI</a> tool from Brew.</p>

<pre><code class="language-shell">brew install cloudfoundry/tap/cf-cli
</code></pre>

<h3>Statping Yaml Config</h3>

<p>You can include multiple environment variables to the configg file. If you include <code>DB_CONN</code> Statping will attempt to automatically connect to that database. View the <a href="https://github.com/statping/statping/wiki/Config-with-.env-File" target="_blank">Full List of Environment Variables</a> to fully customize this config.</p>

<pre><code class="language-yaml">applications:
- name: statping
  docker:
    image: statping/statping
  timeout: 180
  env:
    DB_CONN: sqlite / mysql / postgres
    DB_HOST: ...
    DB_PORT: ...
    DB_DATABASE: ...
    DB_USER: ...
    DB_PASS: ...
</code></pre>

<p>Save this file as <code>manifest.yml</code> and then run: <code>cf push</code></p>

<h3>Advanced Yaml Config</h3>

<pre><code class="language-yaml">applications:
- name: statping
  docker:
    image: statping/statping:((version))
  timeout: 180
  env: ((env))
</code></pre>

<p>Then you can run a command like:</p>

<pre><code class="language-shell">cf push --var version=v0.80.53 --var env='DB_CONN: sqlite'
</code></pre>

<p>Thank you <a href="https://github.com/giner" target="_blank">@giner</a> for creating this documentation.</p>

		</div>
	</div>

	<div class="col-12" v-if='tab === "Development"'>
		<h1 class="h1 mt-5 mb-5 text-muted">Development</h1>
		<span class="spacer"></span>
		<div v-pre>
				<h1>Useful Codebase</h1>

<ul>
<li><p><a href="https://github.com/statping/statping/blob/master/handlers/routes.go" target="_blank">handlers/routes.go</a> - Contains all the HTTP request routes</p></li>

<li><p><a href="https://github.com/statping/statping/blob/master/cmd/commands.go" target="_blank">cmd/commands.go</a> - CLI Commands</p></li>

<li><p><a href="https://github.com/statping/statping/blob/master/types/services/struct.go" target="_blank">types/service/struct.go</a> - Service struct</p></li>

<li><p><a href="https://github.com/statping/statping/blob/master/types/core/struct.go" target="_blank">types/core/struct.go</a> - Core struct</p></li>
</ul>

<h1>Deployment Process</h1>

<p>Statping is a pretty cool server for monitoring your services. The way we deploy might be a little cooler though. Statping is using the most bleeding edge technology to release updates and distribute binary files automatically.</p>

<ol>
<li>Source code commits get pushed to Github</li>
<li><a href="https://github.com/GeertJohan/go.rice" target="_blank">go.rice</a> will compile all the static assets into 1 file (rice-box.go in source)</li>
<li>SASS will generate  a compiled version of the CSS.</li>
<li>Statping Help page is generated by cloning the Wiki repo using <code>go generate</code>.</li>
<li>Travis-CI tests the Golang application.</li>
<li>Travis-CI tests the Statping API using <a href="https://github.com/statping/statping/blob/master/source/tmpl/postman.json" target="_blank">Postman</a>.</li>
<li>If all tests are successful, Travis-CI will compile the binaries using <a href="https://github.com/karalabe/xgo" target="_blank">xgo</a>.</li>
<li>Binaries are code signed using the official <a href="https://github.com/statping/statping/wiki/PGP-Signature" target="_blank">PGP key</a> and compressed.</li>
<li><a href="https://cloud.docker.com/repository/docker/statping/statping/builds" target="_blank">Docker</a> receives a trigger to build for the <code>latest</code> tag.</li>
<li>Travis-CI uploads the <a href="https://github.com/statping/statping/releases" target="_blank">latest release</a> as a tagged version on Github.</li>
<li>Travis-CI updates the <a href="https://github.com/hunterlong/homebrew-statping" target="_blank">homebrew-statping</a> repo with the latest version.</li>
</ol>

<p>And that&rsquo;s it! Statping is ready to be shipped and installed.</p>

		</div>
	</div>

	<div class="col-12" v-if='tab === "Makefile"'>
		<h1 class="h1 mt-5 mb-5 text-muted">Makefile</h1>
		<span class="spacer"></span>
		<div v-pre>
				<p>Here&rsquo;s a simple list of Makefile commands you can run using <code>make</code>. The <a href="https://github.com/statping/statping/blob/master/Makefile" target="_blank">Makefile</a> may change often, so i&rsquo;ll try to keep this Wiki up-to-date.</p>

<p align="center">
<a href="https://github.com/statping/statping/blob/master/Makefile">View Makefile</a>
</p>

<h1>Useful Commands</h1>

<ul>
<li><code>make clean</code> - Cleans the Statping directory for a fresh start</li>
<li><code>make test-deps</code> - Installs required golang libraries for compiling and building</li>
</ul>

<h1>Building</h1>

<ul>
<li><code>make build</code> - Build the golang Statping binary for native OS</li>
<li><code>make compile</code> - Compiles the production frontend Vue</li>
<li><code>make build-all</code> - Cleans, compiles, and builds all binaries</li>
<li><code>make build-linux</code> - Builds Linux binaries</li>
<li><code>make build-mac</code> - Builds MacOSX binaries</li>
<li><code>make build-win</code> - Builds Windows binaries</li>
</ul>

<h1>Testing</h1>

<ul>
<li><code>make test</code> - Cleans, compiles and runs golang unit testing</li>
<li><code>make test-api</code> - Test postman API endpoints</li>
<li><code>make test-cli</code> - Test that run on Github Actions using <code>sqlite</code></li>
</ul>

<h1>Docker</h1>

<ul>
<li><code>make docker-latest</code> - Builds the <code>base</code> and <code>latest</code> Statping images</li>
<li><code>make start</code> - Run the entire Statping stack with multiple databases and Grafana</li>
<li><code>make stop</code> - Stop the entire Statping stack</li>
<li><code>make lite</code> - Run a lite Statping stack that includes Postgres</li>
</ul>

<h1>Required System Libraries</h1>

<ul>
<li>Ubuntu <code>apt-get install build-essential</code></li>
<li>MacOSX <code>sudo xcode-select -switch /Applications/Xcode.app/Contents/Developer</code></li>
<li>Windows <a href="http://gnuwin32.sourceforge.net/packages/make.htm" target="_blank">Install Guide for GNU make utility</a></li>
<li>CentOS/RedHat <code>yum groupinstall &quot;Development Tools&quot;</code></li>
</ul>

		</div>
	</div>

	<div class="col-12" v-if='tab === "GraphQL"'>
		<h1 class="h1 mt-5 mb-5 text-muted">GraphQL</h1>
		<span class="spacer"></span>
		<div v-pre>
				<p>Statping implements the <a href="https://graphql.org/" target="_blank">GraphQL</a> API interface so you can customize the exact data you need within a query. The GraphQL endpoint is at <code>/graphql</code> on your Statping instance and is only available for Authenticated users or while sending the <code>Authorization</code> API Secret.</p>

<p align="center">
<a href="https://github.com/statping/statping/blob/master/handlers/graphql/schema.graphql">View schema.graphql</a>
</p>

<hr>

<h3>Example Queries</h3>

<p>Retrieve the <code>id</code>, <code>name</code>, and <code>public</code> parameters from <code>service</code> #2.</p>

<h6>GraphQL Query:</h6>

<pre><code class="language-graphql">{
  service(id: 2) {
    id
    name
    public
  }
}
</code></pre>

<h6>Response:</h6>

<pre><code class="language-json">{
  &quot;data&quot;: {
    &quot;service&quot;: {
      &quot;id&quot;: 2,
      &quot;name&quot;: &quot;Statping Github&quot;,
      &quot;public&quot;: true
    }
  }
}
</code></pre>

<hr>

<h3>GraphQL Integration Checklist</h3>

<ul>
<li><p>[x] Core (read)</p></li>

<li><p>[x] Services (read)</p></li>

<li><p>[x] Groups (read)</p></li>

<li><p>[x] Users (read)</p></li>

<li><p>[x] Messages (read)</p></li>

<li><p>[x] Failures (read)</p></li>

<li><p>[x] Checkins (read)</p></li>

<li><p>[ ] Core (update)</p></li>

<li><p>[ ] Services (update/create)</p></li>

<li><p>[ ] Groups (update/create)</p></li>

<li><p>[ ] Users (update/create)</p></li>

<li><p>[ ] Messages (update/create)</p></li>

<li><p>[ ] Failures (update/create)</p></li>

<li><p>[ ] Checkins (update/create)</p></li>
</ul>

<hr>

<p>The code for handling GraphQL requests is in <a href="https://github.com/statping/statping/tree/master/handlers/graphql" target="_blank">handlers/graphql</a> and is using <a href="https://github.com/99designs/gqlgen" target="_blank">gqlgen</a> to automatically generate the schema based on the golang structs.</p>

		</div>
	</div>

	<div class="col-12" v-if='tab === "SSL"'>
		<h1 class="h1 mt-5 mb-5 text-muted">SSL</h1>
		<span class="spacer"></span>
		<div v-pre>
				<p>You can run Statping with a valid certificate by including 2 files in the root directory. Although, I personally recommend using NGINX or Apache to serve the SSL and then have the webserver direct traffic to the Statping instance. This guide will show you how to implement SSL onto your Statping server with multiple options.</p>

<h2>LetsEncrypt</h2>

<p>You can have Statping generate SSL certificates with LetsEncrypt by entering the following environment variables. Make sure the domain name is pointing to the address of your Statping instance. Once certificates are generated, they will be stored in the <code>certs</code> folder and will automatically renew 30 days before they expire.</p>

<ul>
<li><code>LETSENCRYPT_ENABLE</code> - Set to <strong>true</strong> to have LetsEncrypt enabled. (defaults to false)</li>
<li><code>LETSENCRYPT_HOST</code> - Domain to generate SSL certificate</li>
<li><code>LETSENCRYPT_EMAIL</code> - Email address that gets sent with the LetsEncrypt Request</li>
</ul>

<h2>SSL Certificate with Statping</h2>

<p>If you have your own SSL certificates, use this method. To run the Statping HTTP server in SSL mode, you must include 2 files in the root directory of your Statping application. The 2 files you must include are:
- <code>server.crt</code> SSL Certificate File
- <code>server.key</code> SSL Certificate Key File</p>

<p>The filenames and extensions must match the exact naming above. If these 2 files are found, Statping will automatically start the HTTP server in SSL mode using your certificates. You can also generate your own SSL certificates, but you will receive a &ldquo;ERR_CERT_AUTHORITY_INVALID&rdquo; error. To generate your own, follow the commands below:</p>

<pre><code class="language-shell">openssl req -new -sha256 -key server.key -out server.csr
openssl x509 -req -sha256 -in server.csr -signkey server.key -out server.crt -days 3650
</code></pre>

<p>This will generate a self signed certificate that you can use for your Statup instance. I recommend using a web server to do SSL termination for your server though.</p>

<h2>Choose a Web Server or Environment</h2>

<p><strong>Choose the environment running the Statping instance.</strong>
- <a href="#docker">Docker</a>
- <a href="#nginx">NGINX</a>
- <a href="#apache">Apache</a></p>

<h2>Docker</h2>

<p>Docker might be the easiest way to get up and running with a SSL certificate. Below is a <code>docker-compose.yml</code> file that will run NGINX, LetEncrypt, and Statping.</p>

<ol>
<li><p>Point your domain or subdomain to the IP address of the Docker server. This would be done on CloudFlare, Route53, or some other DNS provider.</p></li>

<li><p>Replace the <code>docker-compose.yml</code> contents:</p></li>
</ol>

<ul>
<li><code>MY.DOMAIN.COM</code> with the domain you want to use</li>
<li><code>MY@EMAIL.COM</code> with your email address</li>
</ul>

<ol>
<li>Run the docker container by running command <code>docker-compose up -d</code>. Give a little bit of time for LetEncrypt to automatically generate your SSL certificate.</li>
</ol>

<h6><code>docker-compose.yml</code></h6>

<pre><code class="language-yaml">version: '2.3'
services:
  nginx:
    container_name: nginx
    image: jwilder/nginx-proxy
    ports:
      - 0.0.0.0:80:80
      - 0.0.0.0:443:443
    labels:
      - &quot;com.github.jrcs.letsencrypt_nginx_proxy_companion.nginx_proxy&quot;
    networks:
      - internet
    restart: always
    volumes:
      - /var/run/docker.sock:/tmp/docker.sock:ro
      - ./statping/nginx/certs:/etc/nginx/certs:ro
      - ./statping/nginx/vhost:/etc/nginx/vhost.d
      - ./statping/nginx/html:/usr/share/nginx/html:ro
      - ./statping/nginx/dhparam:/etc/nginx/dhparam
    environment:
      DEFAULT_HOST: MY.DOMAIN.COM

  letsencrypt:
    container_name: letsencrypt
    image: jrcs/letsencrypt-nginx-proxy-companion
    networks:
      - internet
    restart: always
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock:ro
      - ./statping/nginx/certs:/etc/nginx/certs
      - ./statping/nginx/vhost:/etc/nginx/vhost.d
      - ./statping/nginx/html:/usr/share/nginx/html
      - ./statping/nginx/dhparam:/etc/nginx/dhparam

  statping:
    container_name: statping
    image: statping/statping:latest
    restart: always
    networks:
      - internet
    depends_on:
      - nginx
    volumes:
      - ./statping/app:/app
    environment:
      VIRTUAL_HOST: MY.DOMAIN.COM
      VIRTUAL_PORT: 8080
      LETSENCRYPT_HOST: MY.DOMAIN.COM
      LETSENCRYPT_EMAIL: MY@EMAIL.COM

networks:
  internet:
    driver: bridge
</code></pre>

<h2>NGINX</h2>

<p>If you already have a NGINX web server running, you just have to add a proxy pass and your SSL certs to the nginx config or as a vhost. By default Statping runs on port 8080, you can change this port by starting server with <code>statping --ip 127.0.0.1 --port 9595</code>.</p>

<ul>
<li>Replace <code>/my/absolute/directory/for/cert/server.crt</code> with SSL certificate file.</li>
<li>Replace <code>/my/absolute/directory/for/key/server.key</code> with SSL key file.</li>
<li>Run <code>service nginx restart</code> and try out https on your domain.</li>
</ul>

<h5>Tutorials</h5>

<ul>
<li><a href="https://docs.nginx.com/nginx/admin-guide/security-controls/terminating-ssl-http/" target="_blank">NGINX Guide</a></li>
<li><a href="https://www.digitalocean.com/community/tutorials/how-to-set-up-nginx-load-balancing-with-ssl-termination" target="_blank">How To Set Up Nginx Load Balancing with SSL Termination</a></li>
</ul>

<h6><code>/etc/nginx/nginx.conf</code></h6>

<pre><code>#user  nobody;
worker_processes  1;
events {
    worker_connections  1024;
}
http {
    include            mime.types;
    default_type       application/octet-stream;
    send_timeout       1800;
    sendfile           on;
    keepalive_timeout  6500;
    server {
        listen       80;
        server_name  localhost;
        location / {
          proxy_pass          http://localhost:8080;
          proxy_set_header    Host             $host;
          proxy_set_header    X-Real-IP        $remote_addr;
          proxy_set_header    X-Forwarded-For  $proxy_add_x_forwarded_for;
          proxy_set_header    X-Client-Verify  SUCCESS;
          proxy_set_header    X-Client-DN      $ssl_client_s_dn;
          proxy_set_header    X-SSL-Subject    $ssl_client_s_dn;
          proxy_set_header    X-SSL-Issuer     $ssl_client_i_dn;
          proxy_read_timeout 1800;
          proxy_connect_timeout 1800;
        }
    }
    # HTTPS server

    server {
        listen       443;
        server_name  localhost;

        ssl                  on;
        ssl_certificate      /my/absolute/directory/for/cert/server.crt;
        ssl_certificate_key  /my/absolute/directory/for/key/server.key;
        ssl_session_timeout  5m;

        ssl_protocols  SSLv2 SSLv3 TLSv1;
        ssl_ciphers  ALL:!ADH:!EXPORT56:RC4+RSA:+HIGH:+MEDIUM:+LOW:+SSLv2:+EXP;
        ssl_prefer_server_ciphers   on;

        location / {
          proxy_pass          http://localhost:8080;
          proxy_set_header    Host             $host;
          proxy_set_header    X-Real-IP        $remote_addr;
          proxy_set_header    X-Forwarded-For  $proxy_add_x_forwarded_for;
          proxy_set_header    X-Client-Verify  SUCCESS;
          proxy_set_header    X-Client-DN      $ssl_client_s_dn;
          proxy_set_header    X-SSL-Subject    $ssl_client_s_dn;
          proxy_set_header    X-SSL-Issuer     $ssl_client_i_dn;
          proxy_read_timeout 1800;
          proxy_connect_timeout 1800;
        }
    }
}
</code></pre>

<h2>Apache</h2>

		</div>
	</div>

	<div class="col-12" v-if='tab === "Config with .env File"'>
		<h1 class="h1 mt-5 mb-5 text-muted">Config with .env File</h1>
		<span class="spacer"></span>
		<div v-pre>
				<p>It may be useful to load your environment using a <code>.env</code> file in the root directory of your Statping server. The .env file will be automatically loaded on startup and will overwrite all values you have in config.yml.</p>

<p>If you have the <code>DB_CONN</code> environment variable set Statping will bypass all values in config.yml and will require you to have the other DB_* variables in place. You can pass in these environment variables without requiring a .env file.</p>

<h2><code>.env</code> File</h2>

<pre><code class="language-bash">DB_CONN=postgres
DB_HOST=0.0.0.0
DB_PORT=5432
DB_USER=root
DB_PASS=password123
DB_DATABASE=root

NAME=Demo
DESCRIPTION=This is an awesome page
DOMAIN=https://domain.com
ADMIN_USER=admin
ADMIN_PASSWORD=admin
ADMIN_EMAIL=info@admin.com
USE_CDN=true
POSTGRES_SSLMODE=false # enable ssl_mode for postgres (To enable use require)
DISABLE_LOGS=false # disable logs from appearing and writing to disk

IS_DOCKER=false
IS_AWS=false
SASS=/usr/local/bin/sass
CMD_FILE=/bin/bash
</code></pre>

<p>This .env file will include additional variables in the future, subscribe to this repo to keep up-to-date with changes and updates.</p>

		</div>
	</div>

	<div class="col-12" v-if='tab === "config.yml"'>
		<h1 class="h1 mt-5 mb-5 text-muted">config.yml</h1>
		<span class="spacer"></span>
		<div v-pre>
				<p>The <code>config.yml</code> file contains all of the Statping database settings and API keys/secrets. When Statping loads, it will attempt to open this file to parse it&rsquo;s core values. You can view the Go type file on <a href="https://github.com/statping/statping/blob/master/types/types.go#L40" target="_blank">types/types.go</a>.</p>

<h3><code>config.yml</code> Parameters</h3>

<ul>
<li><strong>connection</strong>: sqlite (sqlite, mysql, postgres)</li>
<li><strong>host</strong>: 0.0.0.0 (database host)</li>
<li><strong>user</strong>: root (database username)</li>
<li><strong>password</strong>: password123 (database password)</li>
<li><strong>database</strong>: root (Statping database name)</li>
<li><strong>port</strong>: 5432 (database port)</li>
<li><strong>api_key</strong>: 1d31a96a8b3&hellip; (Public API Key)</li>
<li><strong>api_secret</strong>: 18cd0a7373cb&hellip; (Private API Key)</li>
<li><strong>location</strong>: /Users/username/statping_dir (Statping absolute directory)</li>
</ul>

<p>The <code>location</code> setting will contain core files/folders for Statping like, <code>logs</code>, <code>assets</code>, and should contain the <code>config.yml</code> file.</p>

<h2>Using Environment Variables instead of <code>config.yml</code></h2>

<p>You technically don&rsquo;t need this <code>config.yml</code> file to run Statping. You can also run the Statping server by including environment variables when you start the server. You can view all the available variables by viewing the <a href="https://github.com/statping/statping/wiki/Config-with-.env-File" target="_blank">Config with .env</a> Wiki page.</p>

<p><strong>Statping will auto run with environment variables if</strong> <code>DB_CONN</code> <strong>is set!</strong></p>

<h2>Example <code>config.yml</code></h2>

<pre><code class="language-yaml">connection: mysql
host: localhost
user: root
password: password123
database: root
port: 3306
api_secret: apisecretcanbeset
location: /app/statping
</code></pre>

		</div>
	</div>

	<div class="col-12" v-if='tab === "Static HTML Exporter"'>
		<h1 class="h1 mt-5 mb-5 text-muted">Static HTML Exporter</h1>
		<span class="spacer"></span>
		<div v-pre>
				<p>You might have a server that won&rsquo;t allow you to run command that run longer for 60 seconds, or maybe you just want to export your status page to a static HTML file. Using the Statping exporter you can easily do this with 1 command.</p>

<pre><code>statping export
</code></pre>

<h6>&lsquo;index.html&rsquo; is created in current directory with static CDN url&rsquo;s.</h6>

<h2>Push to Github</h2>

<p>Once you have the <code>index.html</code> file, you could technically send it to an FTP server, Email it, Pastebin it, or even push to your Github repo for Status updates directly from repo.</p>

<pre><code class="language-bash">git add index.html
git commit -m &quot;Updated Status Page&quot;
git push -u origin/master
</code></pre>

		</div>
	</div>

	<div class="col-12" v-if='tab === "Prometheus Exporter"'>
		<h1 class="h1 mt-5 mb-5 text-muted">Prometheus Exporter</h1>
		<span class="spacer"></span>
		<div v-pre>
				<p>Statping includes a prometheus exporter so you can have even more monitoring power with your services. The prometheus exporter can be seen on <code>/metrics</code>, simply create another exporter in your prometheus config. Use your Statping API Secret for the Authorization Bearer header, the <code>/metrics</code> URL is dedicated for Prometheus and requires the correct API Secret has <code>Authorization</code> header.</p>

<h1>Grafana Dashboard</h1>

<p>Statping has a <a href="https://grafana.com/dashboards/6950" target="_blank">Grafana Dashboard</a> that you can quickly implement if you&rsquo;ve added your Statping service to Prometheus. Import Dashboard ID: <code>6950</code> into your Grafana dashboard and watch the metrics come in!</p>

<p align="center"><img width="80%" src="https://img.cjx.io/statupgrafana.png"></p>

<h2>Basic Prometheus Exporter</h2>

<p>If you have Statping and the Prometheus server in the same Docker network, you can use the yaml config below.</p>

<pre><code class="language-yaml">scrape_configs:
  - job_name: 'statping'
    scrape_interval: 30s
    bearer_token: 'SECRET API KEY HERE'
    static_configs:
      - targets: ['statping:8080']
</code></pre>

<h2>Remote URL Prometheus Exporter</h2>

<p>This exporter yaml below has <code>scheme: https</code>, which you can remove if you arn&rsquo;t using HTTPS.</p>

<pre><code class="language-yaml">scrape_configs:
  - job_name: 'statping'
    scheme: https
    scrape_interval: 30s
    bearer_token: 'SECRET API KEY HERE'
    static_configs:
      - targets: ['status.mydomain.com']
</code></pre>

<h3><code>/metrics</code> Output</h3>

<pre><code>statping_total_failures 206
statping_total_services 4
statping_service_failures{id=&quot;1&quot; name=&quot;Google&quot;} 0
statping_service_latency{id=&quot;1&quot; name=&quot;Google&quot;} 12
statping_service_online{id=&quot;1&quot; name=&quot;Google&quot;} 1
statping_service_status_code{id=&quot;1&quot; name=&quot;Google&quot;} 200
statping_service_response_length{id=&quot;1&quot; name=&quot;Google&quot;} 10777
statping_service_failures{id=&quot;2&quot; name=&quot;Statping.com&quot;} 0
statping_service_latency{id=&quot;2&quot; name=&quot;Statping.com&quot;} 3
statping_service_online{id=&quot;2&quot; name=&quot;Statping.com&quot;} 1
statping_service_status_code{id=&quot;2&quot; name=&quot;Statping.com&quot;} 200
statping_service_response_length{id=&quot;2&quot; name=&quot;Statping.com&quot;} 2
</code></pre>

		</div>
	</div>

	<div class="col-12" v-if='tab === "PGP Signature"'>
		<h1 class="h1 mt-5 mb-5 text-muted">PGP Signature</h1>
		<span class="spacer"></span>
		<div v-pre>
				<p>You can check if the Statping binary you downloaded is authentic by running a few commands.</p>

<h3>Steps to Authenticate</h3>

<ol>
<li>Download the Statping <code>tar.gz</code> file from <a href="https://github.com/statping/statping/releases/latest" target="_blank">Latest Releases</a> and extract the <code>statping</code> binary and the <code>statup.asc</code> file.</li>
<li>Run command: <code>gpg --verify statping.asc</code></li>
<li>You should see <code>Good signature from &quot;Hunter Long &lt;info@statping.com&gt;&quot; [ultimate]</code>.</li>
</ol>

<h1>Statping Public Key</h1>

<ul>
<li><a href="https://statping.com/statping.gpg" target="_blank">https://statping.com/statping.gpg</a></li>
</ul>

<p>You can also download the key with the command below:</p>

<pre><code>wget https://statping.com/statping.gpg
</code></pre>

<h6><code>statping.gpg</code></h6>

<pre><code>-----BEGIN PGP PUBLIC KEY BLOCK-----

mQINBFwGUYIBEADNsDY4aUOx8EoZuTRFPtjuadJzFRyKtHhw/tLlAnoFACanZPIT
NZoRYvRR5v6lMDXdxsteDbJEOhZ1WDiKIr4OyMahPsyyH6ULzSBKgePUswa0sDef
UnXYzPFQCzqQyQQFbp9AYfDP7dW6dTL9I6qU2NqlJvjxJiiZTAq87SmsLqHiASnI
+ottnQuu6vJQBJz2PFIuaS1c3js/+HBbth9GK5B9YN1BIIyZoFmWKVU9HnJf+aM3
Us6OLjjwYwWzQH38ZV84IjVXyiP9PQVhlCXeHK7XdhPZvnSP1m5Wszj/jowwY6Mz
LgLotfL540X7yOJ7hJTFYLFBOtJdJr/3Ov8SH4HXdPFPVG+UqxsmtmPqUQ9iAxAE
jRFfkAxBvH5Szf2WZdaLnlrrOcOKJIIjZgHqalquBTAhlh5ul0lUVSSPxetwIBlW
60L41k94NJFGDt8xOJ+122mLcywmQ1CzhDfeIKlxl6JDiVHjoRqlQQrqIoNZMV85
rzGfrmbuwv1MXGBJoiNy3330ujOBmhQ9dQVwKpxhBKdjnAgIGM9szbUYxIkGgM1O
U4b1WF3AF/9JOpKJ0LewslpM3BFFYnemGsHXAv3TBPqKidNdwMAiBOtNykGoXF6i
0D6jOW/IB1da0gUA+kr5JdAOwIG7JXKhur2MO7Ncid59DL2N8RePRWj+jwARAQAB
tB9IdW50ZXIgTG9uZyA8aW5mb0BzdGF0cGluZy5jb20+iQJOBBMBCAA4FiEEt21h
+qbbdZRm6D2ZZLnGquLVUngFAlwGUYICGwMFCwkIBwIGFQoJCAsCBBYCAwECHgEC
F4AACgkQZLnGquLVUnizwA//c7vmwTMq/8LYlbo37WM2kDE9AKIrz6VSMq4RhGbC
LikH0X0epa+if79n9BZrVU/Af3aKTn7vu2J4XrvzcdCXtcsR0YmCWML2Y6OSFmhX
w3o6woiFcp+SUWdcM/kithRun+j9sKV4akdgkdBQUdh/RMVln+radz1c6G59iTdh
S+Ip3ObO7Gn5VnrLwxix+W9Jhg8YhDgDGEDt8e1yvjuMRY+WhjHFlwEMoE0kvQL8
QvQH2dGD3dExWAuIL7+0xC0ZGU0PR8vRrq1ukdIsWlDY+42vvhcyPZKFFDTM/QLF
FcCNiPSGhiK/NQq67xnRMFdh0fnqbydWj2atMpacIrheEkOt8db2/UMyDOwlIxgy
KOG8x+yNKiG9LyvW4axRLctN608/+TbvtFo5TVOFJYxJQp4b5uz7LgJAJw7PBvfC
bqx64BH8WGzgyGcAl9unQEtpDuxXoKvP2kbsS7hjvhK0gJgW9llpV4sRJJGApTBc
WtbcS9DBGs3k1aZdA72bxnayD32syVz7czl4+tkRsbQZ4VgJh1yrHIDsdWQXFnYu
EQJfCgX5HvvC13MpDUth0NWCFtWQirY3EFbIgSuhB/D5iXA+Dt1Dq5c1u7wQlUVi
LQCU++oMGrlU3gZrnov5lnBGCEjn0O9bKQm8zmLdEcENFxUZvfPjOIY64YprZxD9
Bv65Ag0EXAZRggEQAMmjHmnvH8SvNJhku/oI96dFKen3bg9xdaFUD1vAuNglCalH
wgXcCZd0RdobYNG46cXTzTQadtHS4hi/UBJ+oy5ZUpIRglW12eTYtqM2G11VbcQi
j6rLITP9NIP+G1xBICSYK4UwmH55BolMEQ/1ZX0a9rESM9stDNglheCCudbMGR/1
ZYnufdEsh0yPwyC/1upZeu8LPWK62pt9mE/gccx77QTeDi5OJcRf1fPbUTCm3vSS
wPPV2AGANodIhostjDymt5vh0tGwc7oUZZLnVdErfuctv7yMgZdiCpYu0jFy1NYf
JgOpZasrcK7/1ozGzsfAo/sSU4kIkMwuWGgqfx5kGRK2CgU4T0i7oI6DMpOX9ZS8
s3+oCWu83X0ofvm5R2CbjiUj2gR6JOhBQbJpCeTkLe+SFcUpnyrr7lG8B8QZHm5N
nBi05V/s63RE3g/6WpR/fWuh+uswe01uqlSx9deW7jT49BL/MdSxwjfwLBLz/hLM
0ld385XAd9bqMjUtp0XhZX2YORx3f/aKY7PYA62baGibb5RdPRw6viEAWU20eb+8
X9Pa7hGmwUeal5lka4SD/TGl7wdY+g4oYP+jtKinH/ZftWA5wHTe3jWT5bdWrT2d
e+0qA0SBkmKIDLpktvtTa19w2nfwBIwJ6fN36ZjYpOn/stxR7aRtnhSqvzxbABEB
AAGJAjYEGAEIACAWIQS3bWH6ptt1lGboPZlkucaq4tVSeAUCXAZRggIbDAAKCRBk
ucaq4tVSeGWmD/9Pg1x6s98zdZCQa2apmUnuoQAQA9Gf2RBBuglCDGsY67wbvdHZ
9wdFRs2QEhl2O3oFmidxthBOBRl9z62nXliLwNn1Lcy/yDfaB8wH6gMm4jn2N/z9
vQXnyIzg8m4PItZ1p5mnY3qH5lpGF8r9Gb7tzK10rqulM2XTDojZOevlEGI6LGw8
FjccXtNquqGZwxzytmKF3T7UBmpmt2qock8N5iJn987m6WeYmbFNc0ii0guHfdtO
zQcItz2ngCdyvfgQPwCAoAv72ysSGhz5KZgAXRrEdcqj6Jw3ivoEUKq1aUrYncXQ
3zC3ED6AjWOGRzjvTZzj22IVacUZ0gqx0x/oldXLOhMB9u6nFXHKj1n9nc0XHMNi
Lp9EuvQgcNLjFZGE9sxh25u9V+OhItfT/aarYEu/Xq0IkUUcdz4GehXth1/Cq1wH
lSUie4nCs7I7OWhqMNClqP7ywElDXsQ66MCgvf01Dh64YUVjJNnyyK0QiYlCx/JQ
Z85hNLtVXZfYqC5BRZlVFp8I8Rs2Qos9YEgn2M22+Rj+RIeD74LZFB7Q4myRvTMB
/P466dFI83KYhwvjBYOP3jPTrV7Ky8poEGifQp2mM294CFIPS7z0z7a8+yMzcsRP
OluFxewsEO0QNDrfFb+0gnjYlnGqOFcZjUMXbDdY5oLSPtXohynuTK1qyQ==
=Xn0G
-----END PGP PUBLIC KEY BLOCK-----
</code></pre>

		</div>
	</div>


<div class="col-12 shadow-md mt-5">
	<div class="text-dim" v-pre>
		<p align="center">
<a href="https://statping.com">Statping.com</a> | <a href="https://demo.statping.com">Demo</a> | <a href="https://hub.docker.com/r/statping/statping">Docker</a> | <a href="https://github.com/statping/statping/wiki/Notifiers">Notifiers</a> | <a href="https://documenter.getpostman.com/view/1898229/RzfiJUd6">API</a>
</p>

	</div>
</div>

<div class="text-center small text-dim" v-pre>
Automatically generated from Statping's Wiki on 2020-09-15 19:09:14.703237 &#43;0000 UTC
</div>

</div>
</template>

<script>
export default {
  name: 'Help',
	data () {
		  return {
			  tab: "Home",
		  }
	  }
}
</script>

<style scoped>
IMG {
	max-width: 80%;
	alignment: center;
	display: block;
}
</style>
