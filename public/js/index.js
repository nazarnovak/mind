(function() {
    'use strict'

    var newCaseB = document.getElementById("newcase");

    newCaseB.addEventListener("click", function() {
        $.post("/api/cases/new", {}, function(id) {
            window.location.href = "/cases/" + id;
        });
    });
})()