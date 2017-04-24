
class Channel extends React.Component{
	onclick(){
		console.log('I was clicked',this.props.name);
	}
	render(){
		return(
			/*Passing the name to the html*/
			<li onClick>{this.onClick.bind(this)} => {this.props.name}</li> 
			)
	}
}

class ChannelList extends React.Component{

	render(){
		return (
			<ul>
			{this.props.channels.map(channel => {
				return (
				<Channel name={channel.name}/>
				)
			}
			)}
			</ul>
			)
	}
}

//Adicionar things to array nao esquecer do bind!!!!
class ChannelForm extends React.Component{
	constructor(props){
		super(props); //has the same properties
		this.state = {}; //initialize the state of the component ITEMS MODIFIED HERE WILL TRIGGER THE RENDER CYCLE
	}
	onChange(e){
		this.setState({
			channelName: e.target.value
		});
			console.log(e.targe.value);
	}
	onSubmit(e)
		{
			let {channelName} = this.state;	
			console.log(channelName);
					this.setState({
			channelName: '' //to make the text disapear after submiting
		});
			this.props.addChannel();
			e.preventDefault();		//prevent browser from sending things
	}
	render(){
		return (
			<form onSubmit={this.onSubmit.bind(this)}>
			<input type='text'
				onChange={this.onChange.bind(this)}
				value={this.state.channelName}
				/>
			</form>
			)
	}
}
//ON THE REACTDOM A UNIQUE THING CAN BE RENDERED (THE TRICK IS RENDER SOMETHING WITH MULTIPLE SUBFILES)

class ChannelSection extends React.Component{
	constructor(props){
		super(props),
		this.state = {
			channels: [
		{ name: 'Embedded Systems'},
		{ name: 'Professores' }]
	};
}
	//create a setter for childs to addChannel
	addChannel(name){
		let {channels}= this.state;
		channels.push({name:name});
		this.setState({
			channels: channels
		});
	}
	render(){
		return (
			<div>
			<ChannelList channels={this.state.channels} />
			<ChannelForm addChannel = {this.addChannel.bind(this)}/>
			</div>
			)
	}
}




ReactDOM.render(<ChannelSection />,
document.getElementById('app'));

/*ReactDOM.render(<Channel name='Embedded Systems'/>,
document.getElementById('app'));*/ 
/*ReactDOM.render(<ChannelList channels ={channels} />,
document.getElementById('app'));*/