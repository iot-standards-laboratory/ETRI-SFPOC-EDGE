window.onload = () => {
  if (!window.WebSocket) {
    alert("No WebSocket!");
    return;
  }
  function connect() {
    let ws = new WebSocket(`ws://${window.location.host}/svc/b68a796a-aceb-4d65-9bbb-6903cba13e56/push/v1`);
    ws.onopen = (e) => {
      console.log("onopen", arguments);
    }

    ws.onclose = () => {
      console.log("onclose", arguments);
    }

    ws.onmessage = function (e) {
      console.log(e.data);
      // addMessage(JSON.parse(e.data));
      console.log(JSON.parse(e.data));
    }
    return ws;
  }

  ws = connect();
}