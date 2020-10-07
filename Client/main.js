var GlobalMessage

function RenderToHtml() {
    var str = document.getElementById("output").innerHTML;
   
    var res = str
              .replace(/&amp;/g, '&')
              .replace(/&quot;/g, '\"')
              .replace(/&#39;/g, '\'')
              .replace(/&lt;/g, '<')
              .replace(/&gt;/g, '>');
   
   
    document.getElementById("output").innerHTML = res;
  }
/*
function interpretMessage(message) {
    //var message = document.getElementById("message-input").value
    
    var url = `http://localhost:8080/interpret?msg=${encodeURIComponent(message)}`

    var request = new XMLHttpRequest()
    request.open("GET", url)
    request.onload = () => {
        var response = JSON.parse(request.responseText) 
        //document.querySelector("#output").querySelector("h3").innerText = response
        
//GlobalMessage = response
        //RenderToHtml();
    }
    
    request.send()
    
}*/

function interpretMessage() {
    var message = document.getElementById("message-input").value
    
    var url = `http://localhost:8080/interpret?msg=${encodeURIComponent(message)}`

    var request = new XMLHttpRequest()
    request.open("GET", url)
    request.onload = () => {
        var response = JSON.parse(request.responseText) 
        document.querySelector("#output").querySelector("h3").innerText = response
        GlobalMessage = response

        RenderToHtml();
    }
    
    request.send()

    //interpretMessage(message);
    
}


function addAcronym() {
    var acronym = document.getElementById("addAcronym-input").value
    var definition = document.getElementById("definition-input").value
    
    var url = `http://localhost:8080/add?acronym=${encodeURIComponent(acronym)}&def=${encodeURIComponent(definition)}`

    var request = new XMLHttpRequest()
    request.open("GET", url)
    request.onload = () => {
        var response = JSON.parse(request.responseText) 
        document.querySelector("#output").querySelector("h3").innerText = response
    }
    
    request.send()
}

function deleteAcronym() {
    var acronym = document.getElementById("delAcronym-input").value

    var url = `http://localhost:8080/del?acronym=${encodeURIComponent(acronym)}`

    var request = new XMLHttpRequest()
    request.open("GET", url)
    request.onload = () => {
        var response = JSON.parse(request.responseText) 
        document.querySelector("#output").querySelector("h3").innerText = response
    }
    
    request.send()
}

function searchAcronyms() {
    var searchAcronym = document.getElementById("searchAcronym-input").value
    
    var url = `http://localhost:8080/search?acronym=${encodeURIComponent(searchAcronym)}`

    var request = new XMLHttpRequest()
    request.open("GET", url)
    request.onload = () => {
        var response = JSON.parse(request.responseText) 
        document.querySelector("#search-output").querySelector("h3").innerText = response
    }

    request.send()
}

function loadPreset(preset) {
    var url = `http://localhost:8080/load?preset=${encodeURIComponent(preset)}`

    var request = new XMLHttpRequest()
    request.open("GET", url)    

    alert("Preset Added!");

    request.send()
}

function textToSpeech() {
    window.open('http://api.voicerss.org/?key=e7e20c508b71469796ca9356a2ed234a&hl=en-us&src=' + GlobalMessage)
}

function downloadString() {
    var blob = new Blob([GlobalMessage], { type: "txt" })
    var a = document.createElement('a')

    a.download = "message"
    a.href = URL.createObjectURL(blob)
    a.dataset.downloadurl = ["txt", a.download, a.href].join(':')
    a.style.display = "none"
    document.body.appendChild(a)
    a.click()
    document.body.removeChild(a)
    setTimeout(function() { URL.revokeObjectURL(a.href) }, 1500)
  }

function textMessage() {
  var phoneNumber = prompt("Enter A Phone Number:", "")

  if (phoneNumber != null) {
    document.getElementById("messageSent").innerHTML =
    "Message Sent to:  " + phoneNumber + "!"
    
    var url = `http://localhost:8080/sendText?phoneNumber=${encodeURIComponent(phoneNumber)}&msg=${encodeURIComponent(GlobalMessage)}`

    var request = new XMLHttpRequest()
    request.open("GET", url)
    request.onload = () => {
        var response = JSON.parse(request.responseText) 
        document.querySelector("#output") = response
    }
    
    request.send()
  }
}

function emailMessage() {
  var emailAddress = prompt("Enter An Email:", "")

  if (emailAddress != null) {
    document.getElementById("messageSent").innerHTML =
    "Email Sent to:  " + emailAddress + "!"
    
    var url = `http://localhost:8080/sendEmail?email=${encodeURIComponent(emailAddress)}&msg=${encodeURIComponent(GlobalMessage)}`

    var request = new XMLHttpRequest()
    request.open("GET", url)
    request.onload = () => {
        var response = JSON.parse(request.responseText) 
        document.querySelector("#output") = response
    }
    
    request.send()
  }
}

function removeAll() {
    var searchAcronym = document.getElementById("searchAcronym-input").value
    var url = `http://localhost:8080/remove?remove=true`

    var request = new XMLHttpRequest()
    request.open("GET", url)
    request.onload = () => {
        var response = JSON.parse(request.responseText) 
        document.querySelector("#remove_message").querySelector("h3").innerText = response
    }

    request.send()
}

function copyToClipboard() {
    var copyText = document.getElementById("merp")
    var textArea = document.createElement("textarea")
    textArea.value = copyText.innerText
    document.body.append(textArea)
    textArea.select()
    document.execCommand("copy");
    document.body.removeChild(textArea)

    alert("Copied the text: " + textArea.value);
}
