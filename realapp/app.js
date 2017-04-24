
class Channel extends React.Component{

	render(){
		return(
			<li>{this.props.name}</li> /*Passing the name to the html*/
			)
	}
}

ReactDOM.render(<Channel name='Embedded Systems'/>,
document.getElementById('app'));