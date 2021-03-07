import React from "react";
import { Block, Button, List, ListInput, Page } from "framework7-react";
import { createCustomer } from "../../services/api";

export default class customer extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      id: "",
      name: "",
      email: "",
      phonenumber: "",
      gstin: "",
    };
  }

  render() {
    return (
      <Page>
        <Block strong>
          <form
            onSubmit={this.onSubmitCustomerClicked}
            action=""
            method="GET"
            className="form-ajax-submit"
          >
            <List>
              <ListInput
                label="Name"
                type="text"
                placeholder="Name"
                value={this.state.name}
                onInput={(e) => {
                  this.setState({ name: e.target.value });
                }}
                required
              ></ListInput>
              <ListInput
                label="Email"
                type="email"
                placeholder="Email address"
                value={this.state.email}
                onInput={(e) => {
                  this.setState({ email: e.target.value });
                }}
                required
              ></ListInput>
              <ListInput
                label="Phone number"
                type="tel"
                placeholder="Phone Number"
                value={this.state.phonenumber}
                onInput={(e) => {
                  this.setState({ phonenumber: e.target.value });
                }}
                required
              ></ListInput>
              <ListInput
                label="GSTIN"
                type="number"
                placeholder="GSTIN"
                value={this.state.gstin}
                onInput={(e) => {
                  this.setState({ gstin: e.target.value });
                }}
                required
              ></ListInput>
            </List>
            <Button fill type="submit">
              Create Customer
            </Button>
          </form>
        </Block>
      </Page>
    );
  }

  showError = (error) => {
    alert(error.message + " Please try again.");
  };

  showSuccess = () => {
    alert("Success");
    this.setState({
      id: "",
      name: "",
      email: "",
      phonenumber: "",
      gstin: "",
    });
  };

  onSubmitCustomerClicked = (e) => {
    e.preventDefault();
    createCustomer(
      this.state.id,
      this.state.name,
      this.state.email,
      this.state.phonenumber,
      this.state.gstin,
      this.showError,
      this.showSuccess
    );
  };
}
