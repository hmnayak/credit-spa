import { Button } from "framework7-react";
import React from "react";

export default class aboutinfo extends React.Component {
  componentDidMount() {
    this.loadDoc("/api/ping", this.myFunction);
  }

  loadDoc(url, cFunction) {
    var xhttp;
    xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function () {
      if (this.readyState == 4 && this.status == 200) {
        cFunction(this);
      }
    };
    xhttp.open("GET", url, true);
    xhttp.send();
  }

  myFunction(xhttp) {
    document.getElementById("aboutinfo").innerHTML = xhttp.responseText;
  }

  render() {
    return (
      <div>
        <div id="aboutinfo" />
      </div>
    );
  }
}
