package utils

import "net/http"

func DiruHttp() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
		    <!DOCTYPE html>
			<html>
				<head>
					<title>Diru</title>
				</head>
				<body>
					<h1>Diru</h1>
					<p>
						If you are seeing this page, it means that you are running Diru on a server. <br/>
						This is purely just to allow for uptime checking for the bot.
					</p>

					<h2>Commands</h2>
					<p>
						<ul>
							<li>@Diru help -- Shows help information</li>
							<li>@Diru info -- Shows various information about the bot</li>
							<li>@Diru <phrase> -- Translates a phrase with the default provider</li>
							<li>@Diru <provider> <phrase> -- Translates a phrase with a specified provider (if enabled)</li>
							<li>@Diru <provider> <lang> <phrase> -- Translates a phrase with a specified provider and target language (if enabled)</li>
						</ul>
					</p>

					<h2>Providers</h2>
					<p>
						<ul>
							<li>gtr -- Google Translate</li>
							<li>dpl -- DeepL</li>
						</ul>
					</p>
				</body>
			</html>
		`))
	})

	http.ListenAndServe(":8080", nil)
}