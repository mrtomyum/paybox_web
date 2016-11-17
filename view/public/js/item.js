$("document").ready(function(){
	active(0);
});
function active_size(id) {
 		$("h1").removeClass("acsize");
		$("#"+id).addClass("acsize");
		//alert('$("#'+id+'").addClass("acsize")');
}

function active(id) {
	if(localStorage.getID!=null){
		console.log(localStorage.getID);
		document.getElementById(localStorage.getID).style.background = "#f3f3f4";
		document.getElementById(localStorage.getID).style.color = "#000";
		localStorage.getID = id;
	}else{
		localStorage.getID = id;
	}	

	$("a").removeClass("active");
	$("#"+id).addClass("active");
	var x = document.getElementsByTagName("LI");
	for(var i = 0;i < x.length; i++){
		x[i].style.background = "#272727";
		if(i!=id){
			//console.log(i);
			x[i].style.borderBottom = "1px dashed #fff";
		}else{
			x[i].style.borderBottom = "0px dashed #fff";
		}
	}
	x[id].style.background = "#272727";	
	document.getElementById(id).style.background = "#272727";
	document.getElementById(id).style.color = "#fff";
	document.getElementById("itemlist").style.background = "#272727";
	document.getElementById("centitem").style.background = "#272727";
}

function showmodal(){
	$("h1").removeClass("acsize");
	$("#s").addClass("acsize");
	$('#myModal').show();
}

