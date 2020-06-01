(function () {
  if (!window.EventSource) {
    alert("No EventSource!");
    return;
  }

  /**
   * @param {string} selector
   * @returns {Element}
   */
  const $ = (selector) => document.querySelector(selector);

  /**
   * @param {string} string
   * @returns {boolean}
   */
  const isBlank = (string) => string == null || string.trim() === "";

  const $chatlog = $("#chat-log");
  const $chatmsg = $("#chat-msg");

  const addmessage = (data) => {
    $chatlog.innerHTML += `<div><span>${data}</span></div>\n`;
  };

  const connect = () => {
    ws = new WebSocket(`ws://${window.location.host}/ws`);
    ws.onopen = (...e) => {
      console.log("onopen", e);
    };
    ws.onclose = (...e) => {
      console.log("onclose", e);
    };
    ws.onmessage = (e) => {
      addmessage(e.data);
    };
  };

  connect();

  let username;
  while (isBlank(username)) {
    username = prompt("what's your name?");
    if (!isBlank(username)) {
      $("#user-name").innerHTML = `<b>${username}</b>`;
    }
  }

  $("#input-form").addEventListener("submit", (e) => {
    if (ws.readyState == ws.OPEN) {
      ws.send(
        JSON.stringify({
          type: "msg",
          data: $chatmsg.value,
        })
      );
    }

    $chatmsg.value = "";
    $chatmsg.focus();

    e.returnValue = false;
  });
})();
