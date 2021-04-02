import React from "react";
import { Block, Button, List, ListInput, Page } from "framework7-react";
import {
  upsertCustomer as upsertCustomer,
  getCustomerApi,
} from "../../services/custapi";

export class CustomerPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      id: "",
      name: "",
      email: "",
      phonenumber: "",
      gstin: "",
      errorMsg: "",
    };
  }

  componentDidMount() {
    let custId = this.props.f7route.params.customerId;
    if (custId != undefined) {
      getCustomerApi(this.props.fetch, custId).then((res) => {
        this.setState({
          id: res.customerid,
          name: res.name,
          email: res.email,
          phonenumber: res.phone,
          gstin: res.gstin,
        });
      });
    }
  }

  submitButton() {
    if (this.props.f7route.params.customerId != undefined) {
      return "Edit Customer";
    } else {
      return "Create Customer";
    }
  }

  render() {
    return (
      <Page>
        <Block strong>
          <form
            onSubmit={this.onSubmitCustomerClicked.bind(this)}
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
                type="text"
                placeholder="GSTIN"
                value={this.state.gstin}
                onInput={(e) => {
                  this.setState({ gstin: e.target.value });
                }}
                required
              ></ListInput>
              <p style={{ color: "red" }}>{this.state.errorMsg}</p>
            </List>
            <Button fill type="submit">
              {this.submitButton()}
            </Button>
          </form>
        </Block>
      </Page>
    );
  }

  showError(error) {
    this.setState({
      errorMsg: error.message,
    });
  }

  showSuccess(status) {
    if (this.props.f7route.params.customerId != undefined) {
      this.props.f7router.navigate("/customers");
    }
    this.props.showNotification(this.state.name);
    this.setState({
      id: "",
      name: "",
      email: "",
      phonenumber: "",
      gstin: "",
    });
  }

  onSubmitCustomerClicked(e) {
    e.preventDefault();
    upsertCustomer(
      this.props.fetch,
      this.state.id,
      this.state.name,
      this.state.email,
      this.state.phonenumber,
      this.state.gstin
    ).then((res) => {
      console.log(res.status);
      if (res.status) {
        this.showSuccess(res.status);
      } else {
        this.showError(res.status);
      }
    });
  }
}
