var selectedChat = "general"

function changeChatRoom() {
    var newchat = document.getElementById("chatroom")
    if(newchat != null && newchat.value != selectedChat){
        console.log(newchat)
    }
    return false;
}

function sendMessage() {
    var newmessage = document.getElementById("message")
    if(newmessage != null){
        console.log(newmessage)
    }
    return false;
}

window.onload = function() {
    document.getElementById("chatroom-selection").onsubmit = changeChatRoom
    document.getElementById("chatroom-message").onsubmit = sendMessage

    if(window["WebSocket"]) {
        console.log("supports websockets")
        conn = new WebSocket("ws://" + document.location.host + "/ws")
    } else {
        alert('Browser does not support websockets')
    }
}
