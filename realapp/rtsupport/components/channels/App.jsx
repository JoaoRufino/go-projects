import React, {Component} from 'react';
import ChannelSection from './ChannelSection.jsx';


class App extends Component{
	constructor(props){
		super(props);
		this.state = {
			channels: []
		};
	}
	addChannel(name){
		let {channels} = this.state;
		channels.push({id:channels.length,name});
		this.setState({channels});
		//TODO : Send to server
	}

	addChannel(activeChannel){
		this.setState({activeChannel});
		//TODO : Send to server
	}
	render(){
		return(
			<ChannelSection
			channels={this.state.channels}
			addChannel={this.addChannel.bind()}
			setChannel={this.state.channels}
			/>
		)
	}


}

export default App