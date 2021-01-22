$(document).ready(function() { 
 
    $('#btn-login').click(function() {  
 
        $(".error").hide();
        var hasError = false;
        var emailReg = /^([\w-\.]+@([\w-]+\.)+[\w-]{2,4})?$/;
        var passreg=/^(?=.*\d)(?=.*[A-Z])(?=.*[a-z]).{8,}$/;
        var passw = $("#Userpass").val();
        var emailaddressVal = $("#UserEmail").val();
        if(emailaddressVal == '') {
            $("#UserEmail").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error">*Enter  your email address.</span>');
            hasError = true;
        }
 
        else if(!emailReg.test(emailaddressVal)) {
            $("#UserEmail").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error">*Enter  valid email address.</span>');
            hasError = true;
        }
 
       
        if(passw == '') {
            $("#Userpass").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error">*Enter your password.</span>');
            hasError = true;
          } 
          else if(!passreg.test(passw)) {
            $("#Userpass").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error">*Enter valid password.</span>');
            hasError = true;
        }
          if(hasError == true) { return false; }
    });

    $('#btn_Register').click(function() {  
  
        $(".error1").hide();
        var hasError = false;
        var emailpat = /^([\w-\.]+@([\w-]+\.)+[\w-]{2,4})?$/;
        var passpat=/^(?=.*\d)(?=.*[A-Z])(?=.*[a-z]).{8,}$/;
        var mobpat=/^([6-9]{1}[0-9]{9})$/;
      
        var name = $("#FirstName").val();
        var mobile = $("#MobileNo").val();
        var email = $("#EmailId").val();
        var pass = $("#Password").val();
        if(email == '') {
            $("#EmailId").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error1">*Enter your email address.</span>');
            hasError = true;
        }
 
        else if(!emailpat.test(email)) {
            $("#EmailId").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error1">*Enter valid email address.</span>');
            hasError = true;
        }
 
       
        if(pass == '') {
            $("#Password").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error1">*Please Set Password.</span>');
            hasError = true;
          } 
          else if(!passpat.test(pass)) {
            $("#Password").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error1">*Password must be in(a-z,A-Z,min 8 digit,one special($,@)).</span>');
            hasError = true;
        }

        if(mobile == '') {
            $("#MobileNo").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error1">*Enter your Mobile no.</span>');
            hasError = true;
          } 
          else if(!mobpat.test(mobile)) {
            $("#MobileNo").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error1">*Mobile number must be 10 digit.</span>');
            hasError = true;
        }

        if(name == '') {
            $("#FirstName").after('<span style="color: rgb(228, 60, 18); font-size: 10pt" class="error1">*Enter your First Name.</span>');
            hasError = true;
          } 

          if(hasError == true) { return false; }
    });

});