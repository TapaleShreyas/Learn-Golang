<!DOCTYPE html>
<html>

<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <link rel="stylesheet" href="https://stackedit.io/style.css" />
</head>

<body class="stackedit">
  <div class="stackedit__html"><h2 id="introduction">Introduction</h2>
<p>Go (also known as the <strong><a href="https://golang.org/">Golang</a></strong>) is an open-source programming language introduced in 2009 and is distributed under the BSD-style license. It was developed by Robert Griesemer, Rob Pike, and Ken Thompson at Google.<br>
In this project, we’ll be learning Golang from basics. This repository will give you the basic programs with some explanation which will help you to understand the concepts. I tried to keep it simple and easily understandable. In order to start typing and running the code, first we need to setup our environment.<br>
<strong>What we need :</strong></p>
<ol>
<li>Windows 10 Home</li>
<li>Go compiler</li>
<li>Visual Studio Code</li>
</ol>
<p>I’m assuming you are using Windows operating system, so let’s starts from how we could install the Go compiler on our Windows operating system. Also how we could configure the Visual Studio Code editor to work with the Go programming language.</p>
<h2 id="installing-go-programming-language">Installing Go Programming Language</h2>
<p>The first thing you need to do is navigate to the golang official site <a href="https://golang.org/dl/"><strong>https://golang.org/dl/</strong></a><br>
This will take you to the downloads page of Go.</p>
<p><img src="https://github.com/TapaleShreyas/Learn-Golang/blob/main/images/download-golang.png" alt="Download Golang"><br>
Download the installer for Windows operating system and start the installation.<br>
The installation process is quite straightforward however if you face any challenges while installing the language do let me know in the comment section.<br>
Once the installation is complete, open your terminal and run the following command.</p>
<pre><code>$ go version
go version go1.16.5 windows/amd64
</code></pre>
<h2 id="add-gopath-environment-variable">Add GOPATH environment variable</h2>
<p>GOPATH is used to specify the root of your Go workspace. By default, the workspace located at %USERPROFILE%/go for Windows.<br>
To configure GOPATH follow the steps mentioned below:</p>
<ol>
<li>Create a new folder in your C drive called a“<em>C:\Users\USER\go</em>”. (in my case)</li>
<li>Open Search box, type environment</li>
<li><img src="https://github.com/TapaleShreyas/Learn-Golang/blob/main/images/search-box.png" alt="Windows Search Box"></li>
<li>Click on <strong>Edit the system Environment variable</strong> , This will open the System Properties</li>
<li>Under Advanced tab, click on <strong>Environment Variables</strong> button at the right bottom.</li>
<li><img src="https://github.com/TapaleShreyas/Learn-Golang/blob/main/images/system-properties.png" alt="System Properties"></li>
<li>Now in the Environment Manager dialog click the new button from the System Variables section.</li>
<li>Set the GOPATH as name and C:\Users\USER\go as its value, and hit the OK button.</li>
<li><img src="https://github.com/TapaleShreyas/Learn-Golang/blob/main/images/go-path-env-variable.png" alt="GOPATH Environment Variable"></li>
</ol>
<p>This will configure the GOPATH on your PC.<br>
To check if the GOPATH has been appropriately configured open run dialog by pressing win + r and type %GOPATH% and hit the OK button. If it takes you to the C:\Users\USER\go directory, It means the configuration is successful.</p>
<h2 id="installing-visual-studio-code">Installing Visual Studio Code</h2>
<p>Visual Studio Code is an open-source lightweight code editor by Microsoft that comes bundled with support for TypeScript, JavaScript, and Node.js. However, you could easily get support for Go Language by installing a handy extension.</p>
<p>Install Visual Studio Code editor, open the link <a href="https://code.visualstudio.com/Download"><strong>https://code.visualstudio.com/Download</strong></a><br>
After the installation process is complete launch the editor and open the extension manager by pressing Ctrl + Shift + x.<br>
In the extension manager, search box type <em>go</em> and hit enter.</p>
<p><img src="https://github.com/TapaleShreyas/Learn-Golang/blob/main/images/vs-go-plugin.png" alt="Visual Studio Golang Extention"><br>
In the search results, you’ll find the Go extension by the GO team at Google. Hit the install button and let the installation process complete.</p>
<p>After installation, open the command palette by pressing Ctrl + Shift + p and run the <em>Go: Install/Update Tools</em> command.</p>
<p><img src="https://github.com/TapaleShreyas/Learn-Golang/blob/main/images/vs-install-update.png" alt="Visual Studio Install/Update Tools"></p>
<p>This will present you with a list of tools you need to install.</p>
<p><img src="https://github.com/TapaleShreyas/Learn-Golang/blob/main/images/vs-install-tools.png" alt="Visual Studio more Tools"></p>
<p>Install everything listed in the dropdown.</p>
<p>Let’s write some code.<br>
Once the installation process is complete create a new hello.go file in your vscode editor and add the following hello world program to the file.</p>
<pre><code>package main
import  "fmt"
func main() {
    fmt.Println("Hello world...")
}

</code></pre>
<p>Now let’s test the autocomplete feature by deleting Println(“Hello world…”) and hitting Ctrl + Space after fmt.</p>
<p><img src="https://github.com/TapaleShreyas/Learn-Golang/blob/main/images/vs-go-editor.png" alt="Visual Studio Go Editor"><br>
You could also try running the program by using the following command in the terminal.</p>
<pre><code>$ go run hello.go
hello world...

</code></pre>
<p>And that’s it!<br>
So, let’s get started with Programming now.</p>
</div>
</body>

</html>
