//Class for creating a Map
// receieves elements and options and creates a MAP

(function(window,google,List) {

	var Mapster = (function() {
		//constructor
		function Mapster(element,options){
			this.gMap= new google.maps.Map(element,options);
			this.markers= List.create(); //guardar todos os markers
		}
		Mapster.prototype = {
		// getter and setters
			zoom: function(level) {
				if(level) {
					this.gMap.setZoom(level);
				} else {
					return this.gMap.getZoom();
				}
			},
			//for events
// https://developers.google.com/maps/documentation/javascript/reference#Maker
//by using opts we can change the thing we apply it to
			_on: function(opts){
				var self = this;
				google.maps.event.addListener(opts.obj, opts.event, function(e){
				opts.callback.call(self,e);
				});
			},
			//add marker
			//in order to keep the other method private
			//not sure if it is the best option
			addMarker: function(opts){
				var marker; //need this to pass to event
				opts.position = {
					lat: opts.lat,
					lng: opts.lng
				}
				marker=this._createMarker(opts);
				//this.markers.add(marker);
				if(opts.event)
				{
					this._on({
						obj: marker,
						event: opts.event.name,
						callback: opts.event.callback
					});
				}
				if(opts.content) {
					var infoWindow; // to pass to other function
					this._on({
						obj: marker,
						event: 'mouseover',
						callback: function(){
							infoWindow= new google.maps.InfoWindow({
								content: opts.content
							});
							infoWindow.open(this.gMap,marker);
						}

					})
					this._on({
						obj: marker,
						event: 'mouseout',
						callback: function(){
							infoWindow.close()
						}
					});
					
					}
				return marker; //for info window
			},
			findBy: function(callback){
				return this.markers.find(callback);
			},

			//createMarker
			//no need for new parameters, just add things to opts
			_createMarker: function(opts){
				opts.map = this.gMap;
				return new google.maps.Marker(opts);
			}

		};
		return Mapster;
	}());
	//self creation function
	Mapster.create=function(element,options){
		return new Mapster(element,options);
	};
	//connect to the window
	window.Mapster = Mapster;

}(window,google,List))