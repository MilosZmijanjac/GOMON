document.addEventListener('DOMContentLoaded', function() {
    const token = localStorage.getItem('jwtToken');
    if (!token) {
        window.location.href = '../login/index.html';
        return;
    }
    const user = localStorage.getItem('selectedUser');
    if (!user) {
        window.location.href = '../users/users.html';
        return;
    }
    var userObj=JSON.parse(user);
    console.log(userObj)
    const form = document.getElementById('editUserForm');
    const submitButton = document.getElementById('submitButton');
    const cancelButton = document.getElementById('cancelButton');
    const formData1 = new FormData(form);
    document.getElementById('username').placeholder = userObj.Username;

    submitButton.addEventListener('click', function(event) {
        if (!form.checkValidity()) {
            console.log("invalid")
            event.preventDefault();
            event.stopPropagation();
            form.reportValidity(); 
            return;
        }
        console.log("submit")
        const formData = new FormData(document.getElementById('editUserForm'));
        console.log(formData)
        const data = {
            password: formData.get('password'),
            role: Number(formData.get('role')),
            isActive:Boolean( formData.get('activity'))
        };

        fetch('http://localhost:8000/user/update', { 
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${localStorage.getItem('jwtToken')}` },
            body: JSON.stringify(data)
        })
        .then(data => {
                alert('User edited successfully!');
                form.reset()
                window.location.href = '../users/users.html'; 
        })
        .catch(error => {
            console.error('Error:', error);
            form.reset()
            alert('An error occurred while editing the user.');
        });
    });

    cancelButton.addEventListener('click', function() {
        window.location.href = '../users/users.html'; 
    });
});
