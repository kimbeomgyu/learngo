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

  let username;
  while (isBlank(username)) {
    username = prompt("what's your name?");
    if (!isBlank(username)) {
      $("#user-name").innerHTML = `<b>${username}</b>`;
    }
  }

  $("#input-form").addEventListener("submit", (e) => {
    if ($chatmsg.value == "") {
      e.returnValue = false;
      return;
    }

    const formData = new FormData();
    formData.append("msg", $chatmsg.value);
    formData.append("name", username);

    fetch("/messages", {
      method: "POST",
      body: formData,
    });

    $chatmsg.value = "";
    $chatmsg.focus();

    e.returnValue = false;
  });

  /**
   * @param {{name?:string,msg:string}} data
   * @returns {void}
   */
  const addMessage = (data) => {
    let text = "";
    if (!isBlank(data.name)) {
      text = `<strong>${data.name}: </strong>`;
    }
    text += data.msg;

    const div = document.createElement("div");
    div.innerHTML = `<span>${text}</span>`;
    $chatlog.prepend(div);
  };

  // EventSource
  const es = new EventSource("/stream");

  // user가 생성되었을 때
  es.onopen = () => {
    const formData = new FormData();
    formData.append("name", username);
    fetch("/users", { method: "POST", body: formData });
  };

  // message를 입력받을 때
  es.onmessage = (e) => {
    const msg = JSON.parse(e.data);
    addMessage(msg);
  };

  // user가 떠날 때
  window.onbeforeunload = () => {
    fetch("/users?username=" + username, { method: "DELETE" });
    es.close();
  };
})();
