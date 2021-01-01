import React, { useState } from "react";
import { App, View, Navbar, f7 } from "framework7-react";
import { user, setNavigate } from "./auth";
import routes from "./routes";

const rootPath = window.location.pathname.replace(/\/+$/, "");

export default class Container extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      isLoading: false,
    };
  }

  componentDidMount() {
    console.log(this);
  }

  render() {
    return (
      <App
        name="Credit"
        theme="auto"
        id="treeples.credit"
        routes={routes(this.setLoading)}
      >
        <Navbar title="Credit">{this.loading()}</Navbar>
        <View
          main
          url={rootPath}
          browserHistory
          browserHistorySeparator=""
          browserHistoryRoot=""
          animate={false}
        />
      </App>
    );
  }

  setLoading(isLoading) {
    this.setState({ isLoading: isLoading });
  }

  loading() {
    if (this.state.isLoading) {
      return <div>Loading...</div>;
    }
  }
}

// export default (props) => {
//   const [isLoading, setLoading] = useState(false);

//   const Loading = () => {
//     if (isLoading) {
//       return <div>Loading...</div>;
//     }
//   };

//   const router = Framework7.instance.views.main.router;

//   console.log(this.$f7);
//   // setNavigate(props.f7router.navigate);

//   return (
//     <App name="Credit" theme="auto" id="treeples.credit" routes={routes}>
//       <Navbar title="Credit">{Loading()}</Navbar>
//       <View
//         main
//         url={rootPath}
//         browserHistory
//         browserHistorySeparator=""
//         browserHistoryRoot=""
//         animate={false}
//       />
//     </App>
//   );
// };
