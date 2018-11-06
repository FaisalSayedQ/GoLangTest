var CurrFromText;
var CurrFromAmount;
var CurrToText;
var CurrToAmount;

$('#time').html(new Date());

function ajaxSuccess (response) {
    $('main').html(response);
    $('#time').html(new Date());
    $('#CurrFromText').val(CurrFromText);
    $('#CurrToText').val(CurrToText);

}

function callAjax () {
    CurrFromText = $('#CurrFromText').val();
    CurrFromAmount = parseInt($('#CurrFromAmount').val(), 10) || 0;
    CurrToText = $('#CurrToText').val();
    CurrToAmount = parseInt($('#CurrToAmount').val(), 10) || 0;

    if ((CurrFromText === 'Currency' || CurrToText === 'Currency') || (!CurrFromAmount && !CurrToAmount)) {
        return;
    }

    var url =  '/convert'
        + '/' + $('#CurrFromText').val()
        + '/' + CurrFromAmount 
        + '/' + $('#CurrToText').val() 
        + '/' + CurrToAmount;

    console.log(url);
    
    $.ajax({
      url: url,
      success: ajaxSuccess
    });
}
$('body').on('change', '#CurrFromText', function () {
    callAjax();
});

$('body').on('change', '#CurrToText', function () {
    callAjax();
});

$('body').on('blur', '#CurrFromAmount', function () {
    callAjax();
});

$('body').on('blur', '#CurrToAmount', function () {
    callAjax();
});