function addcart(id) {
     var product_id = id
     $.ajax({
          url: "/addtoorder",
          type: "POST",
          data: {
               product_id: product_id,
          },
          dataType: 'json',
          success: function (results) {


               $('#snackbar').html(results.msg);
               $('#snackbar1').html(results.count);
          }
     });
     var x = document.getElementById("snackbar");
     x.className = "show";
     setTimeout(function () { x.className = x.className.replace("show", ""); }, 3000)
}

function addwish(id) {


     var clname = $('#add_wish' + id).attr('class');

     if (clname === "flex") {
          $('#add_wish' + id).removeAttr('href');

          $.ajax({
               url: "/addwishlists",
               type: "POST",
               data: {
                    product_id: id,
               },
               dataType: 'json',
               success: function (result) {

                    $('#snackbar').html(result);
               }
          });

          $('#add_wish' + id).addClass("added-wish")

          var x = document.getElementById("snackbar");
          x.className = "show";
          setTimeout(function () { x.className = x.className.replace("show", ""); }, 3000)

     } else {
          window.location = "/wishlist";

     }

}

// function qtyplus(id) {
//      var qty = $('.quantity' + id).text();

//      if (qty >= 1 && 6>qty) {
//      $.ajax({
//           url: "/qtyplus",
//           type: "POST",
//           data: {
//                product_id: id,
//           },
//           dataType: 'json',
//           success: function (result) {

//                $('.quantity' + id).html(result);
//           }
//      });
// }
// }

// function qtyminus(id) {
//      var qty = $('.quantity' + id).text();

//      if (qty > 1 && 6>qty) {
//           $.ajax({
//                url: "/qtyminus",
//                type: "POST",
//                data: {
//                     product_id: id,
//                },
//                dataType: 'json',
//                success: function (result) {

//                     $('.quantity' + id).html(result);
//                }
//           });
//      }
// }




