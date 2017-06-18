
$(function(){
$('#text').keydown(function (e){
    if(e.keyCode == 13){
        alert('you pressed enter ^_^');
    }else {
	alert('no enter');
}
})
});