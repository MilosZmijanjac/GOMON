document.addEventListener('DOMContentLoaded', async () => {
    const logoutButton = document.getElementById('logout');
    const addNewDevice=document.getElementById('addNewDevice');
    if (logoutButton) {
        logoutButton.addEventListener('click', function() {
            localStorage.removeItem('jwtToken');
            localStorage.removeItem('user');
            localStorage.removeItem('selectedUser');
            localStorage.removeItem('selectedDevice');
            window.location.href = '../login/index.html'; 
        });
    }
    if (addNewDevice) {
        addNewDevice.addEventListener('click', function() {
            window.location.href = '../add-device/add-device.html'; 
        });
    }
})
document.addEventListener('DOMContentLoaded', async () => {
    const token = localStorage.getItem('jwtToken');
    if (!token) {
        window.location.href = '../login/index.html';
        return;
    }
    const username = localStorage.getItem('user');
    const tableBody = document.getElementById('devicesTable').getElementsByTagName('tbody')[0];
    const isAdmin=true;
    
        const response = await fetch('http://localhost:8000/device/list', {
            method: 'POST',
            headers: {
                'Authorization': `Bearer ${token}`
            },
            body:JSON.stringify({ username,isAdmin })
        });

        if (response.ok) {
            const devices = await response.json();
            devices.devices.forEach(device => {
                const row = tableBody.insertRow();

                row.insertCell(0).textContent = device.DeviceID;
                row.insertCell(1).textContent = device.Name;
                row.insertCell(2).textContent = device.Status;
                row.insertCell(3).textContent = device.IP;
                row.insertCell(4).textContent = device.FirmwareVersion;
                row.insertCell(5).textContent = device.Location;
                row.insertCell(6).textContent = device.OS;

                  const select = document.createElement('select');
                  device.UserIds.forEach(userId => {
                      const option = document.createElement('option');
                      option.value = userId;
                      option.textContent = userId;
                      select.appendChild(option);
                  });
  
                  const cell = row.insertCell(7);
                  cell.appendChild(select);
                  row.addEventListener('click', function() {
                     localStorage.setItem('selectedDevice',JSON.stringify(device))
                     window.location.href = "../device-details/device-details.html";
                });
            });
        } else {
            console.error('Error fetching devices');
            window.location.href = '../login/index.html';
        }

});