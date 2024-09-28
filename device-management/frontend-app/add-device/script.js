document.addEventListener('DOMContentLoaded', function() {
    const token = localStorage.getItem('jwtToken');
    if (!token) {
        window.location.href = '../login/index.html';
        return;
    }
    const form = document.getElementById('addDeviceForm');
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
        const formData = new FormData(document.getElementById('addDeviceForm'));
        console.log(formData)
        const data = {
            deviceId: formData.get('deviceId'),
            name: formData.get('name'),
            status: formData.get('status'),
            ip: formData.get('ip'),
            firmwareVersion: formData.get('firmwareVersion'),
            location: formData.get('location'),
            os: formData.get('os'),
            userIds: formData.get('userIds').split(',').map(id => id.trim()) 
        };

        fetch('http://localhost:8000/device/register', { 
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': `Bearer ${localStorage.getItem('jwtToken')}` },
            body: JSON.stringify(data)
        })
        .then(data => {
                alert('Device added successfully!');
                form.reset()
                window.location.href = '../devices/devices.html'; 
        })
        .catch(error => {
            console.error('Error:', error);
            form.reset()
            alert('An error occurred while adding the device.');
        });
    });

    cancelButton.addEventListener('click', function() {
        window.location.href = '../devices/devices.html'; 
    });
});
