<!DOCTYPE html>
<html>
<head>
  <meta charset="UTF-8">
</head>
<body>
  <h1>CLI Application for IP and Server Name Lookup</h1>
  <p>This is a command-line interface (CLI) application written in Go (version 1.20) that allows you to search for IP addresses and server names. It utilizes the <code>net</code> package to perform DNS lookups.</p>
  <h2>Usage</h2>
  <p>To use this application, follow the steps below:</p>
  <ol>
    <li>Clone the repository:</li>
    <pre><code>git clone https://github.com/psnavega/netSleuth.git</code></pre>
    <li>Build the application:</li>
    <pre><code>go build</code></pre>
    <li>Run the application:</li>
    <pre><code>./app [command] [flags]</code></pre>
  </ol>
  <p>The available commands are:</p>
  <ul>
    <li><code>ip</code>: Search for IP addresses on the internet.</li>
    <li><code>server</code>: Find the names of servers.</li>
  </ul>
  <p>The following flags can be used with both commands:</p>
  <ul>
    <li><code>--host</code>: Specify the host for the lookup (default: <code>globo.com.br</code>).</li>
  </ul>
  <h2>Examples</h2>
  <p>Here are some example commands:</p>
  <ul>
    <li>Search for IP addresses:</li>
    <pre><code>./app ip --host example.com</code></pre>
    <p>This will search for IP addresses associated with the host <code>example.com</code>.</p>
    <li>Find server names:</li>
    <pre><code>./app server --host example.com</code></pre>
    <p>This will find the names of servers associated with the host <code>example.com</code>.</p>
  </ul>
  <h2>Dependencies</h2>
  <p>This application uses the <code>urfave/cli</code> package for building command-line applications. You can install it using the following command:</p>
  <pre><code>go get -u github.com/urfave/cli</code></pre>
  <h2>License</h2>
  <p><a href="https://github.com/psnavega/netSleuth/blob/master/LICENSE">MIT License</a></p>
  <p>Feel free to modify and enhance this application as per your requirements. If you encounter any issues or have suggestions for improvements, please create an issue on the GitHub repository.</p>
</body>
</html>