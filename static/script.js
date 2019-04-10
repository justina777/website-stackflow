
$("a.show-detail").on("click", function(){
    console.log('show-detail');
    console.log($(this).data('id'));
    $("#divItem").show();
    $("#divItem #qid").html($(this).data('id'));
});