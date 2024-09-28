document.addEventListener('DOMContentLoaded', function() {
    const token = localStorage.getItem('jwtToken');
    if (!token) {
        window.location.href = '../login/index.html';
        return;
    }
    const form = document.getElementById('addUserForm');
    const submitButton = document.getElementById('submitButton');
    const cancelButton = document.getElementById('cancelButton');

    submitButton.addEventListener('click', function(event) {
        if (!form.checkValidity()) {
            console.log("invalid")
            event.preventDefault();
            event.stopPropagation();
            form.reportValidity(); 
            return;
        }
        console.log("submit")
        const formData = new FormData(document.getElementById('addUserForm'));
        console.log(formData)
        const data = {
            username: formData.get('username'),
            password: formData.get('password'),
            role: Number(formData.get('role')),
            isActive:Boolean( formData.get('activity'))
        };

        fetch('http://localhost:8000/user/register', { 
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${localStorage.getItem('jwtToken')}` },
            body: JSON.stringify(data)
        })
        .then(data => {
                alert('User added successfully!');
                form.reset()
                window.location.href = '../users/users.html'; 
        })
        .catch(error => {
            console.error('Error:', error);
            form.reset()
            alert('An error occurred while adding the user.');
        });
    });

    cancelButton.addEventListener('click', function() {
        window.location.href = '../users/users.html'; 
    });
});
