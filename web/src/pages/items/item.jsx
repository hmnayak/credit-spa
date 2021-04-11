import React from "react";
import { Block, Button, List, ListInput, Page } from "framework7-react";
import {upsertItem, getItemApi} from "../../services/itemapi";

export class ItemPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      id: "",
      name: "",
      type: "",
      hsn: "",
      sac: "",
      gst: "",
      igst: ""
    };
  }

  componentDidMount() {
    let custId = this.props.f7route.params.itemId;
    if (custId != undefined) {
      getItemApi(this.props.fetch, custId).then(async (res) => {
        const item = await res.json();
        this.setState({
          id: item.itemid,
          name: item.name,
          type: item.type,
          hsn: item.hsn,
          sac: item.sac,
          gst: item.gst,
          igst: item.igst
        });
      });
    }
  }

  submitButton() {
    if (this.props.f7route.params.itemId != undefined) {
      return "Edit item";
    } else {
      return "Create item";
    }
  }

  render() {
    return (
      <Page>
        <Block strong>
          <form
            onSubmit={this.onSubmititemClicked.bind(this)}
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
                onInput={e => this.setState({ name: e.target.value })}
                required
              ></ListInput>
              <ListInput
                label="Type"
                type="text"
                placeholder="Type"
                value={this.state.type}
                onInput={e => this.setState({ type: e.target.value })}
              ></ListInput>
              <ListInput
                label="HSN"
                type="number"
                placeholder="HSN"
                value={this.state.hsn}
                onInput={e => this.setState({ hsn: e.target.value })}
              ></ListInput>
              <ListInput
                label="SAC"
                type="number"
                placeholder="SAC"
                value={this.state.sac}
                onInput={e => this.setState({ sac: e.target.value })}
              ></ListInput>
              <ListInput
                label="GST"
                type="number"
                placeholder="GST"
                value={this.state.gst}
                onInput={e => this.setState({ gst: e.target.value })}
              ></ListInput>
              <ListInput
                label="IGST"
                type="number"
                placeholder="IGST"
                value={this.state.igst}
                onInput={e => this.setState({ igst: e.target.value })}
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
    if (this.props.f7route.params.itemId !== undefined) {
      this.props.f7router.navigate("/items/");
    }
    this.props.showNotification(this.state.name);
    if (this.props.f7route.params.itemId === undefined) {
      this.setState({
        id: "",
        name: "",
        type: "",
        hsn: "",
        sac: "",
        gst: "",
        igst: ""
      });
    }
  }

  onSubmititemClicked(e) {
    e.preventDefault();
    upsertItem(
      this.props.fetch,
      this.state.id,
      this.state.name,
      this.state.type,
      this.state.hsn,
      this.state.sac,
      this.state.gst,
      this.state.igst
    ).then((res) => {
      if (res.ok) {
        this.showSuccess(res.status);
      } else {
        this.showError(res.status);
      }
    });
  }
}
