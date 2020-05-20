(function () {
  if (!window.EventSource) {
    alert("No EventSourc!");
    return;
  }
  const $ = function (selector) {
    return document.querySelector(selector);
  };

  var $chatlog = $("#chat-log");
  var $chatmsg = $("#chat-msg");

  var isBlank = function (string) {
    return string == null || string.trim() === "";
  };
  var username;
  while (isBlank(username)) {
    username = prompt("what's your name?");
    if (!isBlank(username)) {
      $("#user-name").innerHTML = `<b>${username}</b>`;
    }
  }
  $("#input-form").addEventListener("submit", async function (e) {
    const formData = new FormData();
    formData.append("msg", $chatmsg.value);
    formData.append("name", username);

    const rawResponse = await fetch("/messages", {
      method: "POST",
      body: formData,
    });

    const content = rawResponse;
    console.log(content);
    $chatmsg.innerText = "";
    $chatmsg.focus();
  });
})();
