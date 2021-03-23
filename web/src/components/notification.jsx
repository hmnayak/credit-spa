import React from "react";
import "../../css/notification.css";

export class NotificationMsg extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      message: this.props.successMsg,
    };
  }

  componentDidUpdate() {
    if (this.state.message) {
      this.setExpiry();
    }
    this.setState({
      message: this.props.successMsg,
    });
  }

  componentDidMount() {
    this.setExpiry();
  }

  setExpiry() {
    setTimeout(() => {
      this.setState({
        message: "",
      });
    }, 3000);
  }

  shouldComponentUpdate(nextProps, nextState) {
    return nextProps.successMsg != nextState.message;
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
