document.addEventListener('DOMContentLoaded', function (){
    const button = document.querySelector('.addContact');
    button.addEventListener('click', goToForm);

    function goToForm() {
        window.location.href = "addconctact.html";
    }
});

document.addEventListener('DOMContentLoaded', function () {
    const addButton = document.querySelector('.button input');
  
    const formInputs = document.querySelectorAll('.user .input input');
    formInputs.forEach(function (input) {
      input.addEventListener('input', function () {
        validateForm();
      });
});
  
    function validateForm() {
      const firstname = document.getElementById('firstname').value;

      const address = document.getElementById('address').value;
      const email = document.getElementById('email').value;
      const phone = document.getElementById('phone').value;
  
      let isValid = true;
  
      //firstname
      if (firstname.trim() === '') {
        document.getElementById('firstname').classList.add('error');
        isValid = false;
      } else {
        document.getElementById('firstname').classList.remove('error');
      }
      //address
      if (address.trim() === '') {
        document.getElementById('address').classList.add('error');
        isValid = false;
      } else {
        document.getElementById('address').classList.remove('error');
      }
  
      //email
      const emailRegex = /^\S+@\S+\.\S+$/;
      if (!emailRegex.test(email)) {
        document.getElementById('email').classList.add('error');
        isValid = false;
      } else {
        document.getElementById('email').classList.remove('error');
      }
  
      // phone number
      const phoneRegex = /^\d{10}$/;
      if (!phoneRegex.test(phone)) {
        document.getElementById('phone').classList.add('error');
        isValid = false;
      } else {
        document.getElementById('phone').classList.remove('error');
      }
  
      //btn
      if (firstname.trim() === '' || lastname.trim() === '' || address.trim() === '' || email.trim() === '' || phone.trim() === '') {
        addButton.disabled = true;
      } else {
        addButton.disabled = false;
      }
  
      return isValid;
    }
  });

    
   

