// var r1 = document.getElementById('r1');
// var message = document.getElementById('d1');
// r1.addEventListener('click', function () {
//     if(r1.checked) {
//         message.style.display = 'block';
//     } else {
//         message.style.display = 'none';
//     }
// });

// $(document).ready(function (){
//     $("#r1").click(function(){
//       let check_box = $(this);
//       check_val = check_box.val();
//       alert(check_val);
//     });
//   })


$(function () {
    $("#r1").click(function () {
      if($(this).is(":checked")) {
        $("#d1").show();
      } else {
        $("#d1").hide();
      }
    })
  })