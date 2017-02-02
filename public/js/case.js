(function() {
    'use strict'

    var socket = glue();

    var messages = [];

    function getMessages(since) {
        var params = {};
        var messages = {};

        if(since !== undefined) {
            params.since = since;
        }

        $.get("/api/cases/" + caseId + "/events", params, function(data) {
            appendMessages(data)
        });
    }

    function appendMessages(msgs) {
        var msgLog = document.getElementById("log");

        messages = messages.concat(msgs);

        $.each(msgs, function( key, msg) {
            msgLog.value += "[" + msg.Created + "] " + msg.UserName + ": "
                + msg.Content + "\n";
        });
    }

    getMessages();

    socket.send("sub case:" + caseId);

    socket.onMessage(function(data) {
        data = data.split(":");
        if (data[0] === "message") {
            getMessages(messages[messages.length - 1].Created);
        }
    });

    function sendMessage() {
        var msg = msgI.value;
        var data = {
                userid: userId,
                username: userName,
                msg: msg
        };

        $.post("/api/cases/" + caseId + "/events", JSON.stringify(data));

        msgI.value = "";
    }

    var msgI = document.getElementById("msg");
    var sendB = document.getElementById("send");

    window.addEventListener("keypress", function(e) {
        if(e.keyCode == 13) {
            sendMessage();
        }
    })

    sendB.addEventListener("click", function() {
        sendMessage();
    });
})()