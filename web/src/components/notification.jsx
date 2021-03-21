import React from "react";
import "../../css/notification.css";

export class NotificationMsg extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      message: this.props.message,
    };
  }

  componentDidMount() {
    setTimeout(() => {
      this.setState({
        message: "",
      });
    }, 3000);
  }

  render() {
    if (this.state.message) {
      return (
        <div className="notification-wrapper">
          <div className="notification-item">
            <p> {this.state.message}</p>
          </div>
        </div>
      );
    } else return <div></div>;
  }
}
