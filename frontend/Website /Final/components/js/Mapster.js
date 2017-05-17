//Class for creating a Map
// receieves elements and options and creates a MAP

(function(window,google,List) {

	var Mapster = (function() {
		var ws;
		//constructor
		function Mapster(element,opts){
			this.gMap= new google.maps.Map(element,opts);
			this.markers= List.create(); //guardar todos os markers
			if(opts.cluster) {
			ws = new WebSocket("ws://echo.websocket.org/")
			this.markerClusterer = new MarkerClusterer(this.gMap,[],opts.cluster.options);//we need to access it inside functions
			}
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
				opts.callback.call(self,e,opts.obj);
				});
			},
			//add marker
			//in order to keep the other method private
			//not sure if it is the best option
			addMarker: function(opts){
				var marker, //need this to pass to event
				self = this;
				opts.position = {
					lat: opts.lat,
					lng: opts.lng
				}
				marker=this._createMarker(opts);
				if(this.markerClusterer){
				this.markerClusterer.addMarker(marker);
				}
				this._addMarker(marker);
				//this.markers.add(marker);
				if(opts.events){
					this._attachEvents(marker,opts.events);
				}
				if(opts.content) {
					var infoWindow= new google.maps.InfoWindow({
								content: opts.content
							});
					 // to pass to other function
					this._on({
						obj: marker,
						event: 'mouseover',
						callback: function(){
							infoWindow.open(this.gMap,marker);
						}
					}),
					this._on({
						obj: marker,
						event: 'mouseout',
						callback: function(){
							infoWindow.close();
						}
					})
				}
				return marker; //for info window
			},
			_attachEvents: function(obj, events){
				var self = this;
				events.forEach(function(event){
					self._on({
						obj: obj,
						event: event.name,
						callback: event.callback
						});
					});
			},
			_addMarker: function(callback){
				this.markers.find(callback);
			},

			findBy: function(callback){
				this.markers.find(callback);
			},
			removeBy:function(callback){
				var self = this;
				self.markers.find(callback,function(markers){
					markers.forEach(function(marker){
						if(self.markerClusterer){
							self.markerClusterer.removeMarker(marker);
						} else{
							marker.setMap(null);
						}
					});
				});
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
	Mapster.create=function(element,opts){
		return new Mapster(element,opts);
	};
	//connect to the window
	window.Mapster = Mapster;

}(window,google,List))