import React from "react";

export default class ErrorBoundary extends React.Component {
    constructor(props) {
      super(props);
      this.state = { error: null };
    }
  
    static getDerivedStateFromError(error) {
      return { error: error };
    }
  
    componentDidCatch(error, errorInfo) {
      console.log('componentDidCatch', error, errorInfo);
    }
  
    render() {
      if (this.state.error) {
        return (
          <div>
            Unhandled Error. Check console for details
            <div>{this.state.error.message}</div>
          </div>
        );
      }
  
      return this.props.children; 
    }
}