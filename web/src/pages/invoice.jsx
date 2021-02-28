import React from "react";
import { Block, Button, List, ListInput, Page } from "framework7-react";
import { createInvoice } from "../services/api";

export default class invoice extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
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
            onSubmit={this.onSubmitInvoiceClicked}
            action=""
            method="GET"
            className="form-ajax-submit"
          >
            <List>
              <ListInput
                label="Username"
                type="text"
                placeholder="User name"
                value={this.state.name}
                onInput={(e) => {
                  this.setState({ name: e.target.value });
                }}
                required
              ></ListInput>
              <ListInput
                label="Login with your email"
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
              Submit Invoice
            </Button>
          </form>
        </Block>
      </Page>
    );
  }

  showError = (error) => {
    console.error("Failed to login", error);
    alert(error.message + " Please try again.");
  };

  onSubmitInvoiceClicked = (e) => {
    e.preventDefault();

    createInvoice(
      this.state.name,
      this.state.email,
      this.state.phonenumber,
      this.state.gstin,
      this.showError
    );
  };
}
