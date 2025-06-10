function toggleEdit(section) {
    const displaySection = document.getElementById(section + 'Display');
    const editSection = document.getElementById(section + 'Edit');
    const parentInfoItem = displaySection.closest('.info-item');

    if (displaySection && editSection && parentInfoItem) {
        parentInfoItem.style.display = 'none';
        editSection.style.display = 'flex';
    }
}

function cancelEdit(section) {
    const displaySection = document.getElementById(section + 'Display');
    const editSection = document.getElementById(section + 'Edit');
    const parentInfoItem = displaySection.closest('.info-item');

    if (displaySection && editSection && parentInfoItem) {
        parentInfoItem.style.display = 'flex';
        editSection.style.display = 'none';
    }
}

function saveUsername() {
    const newUsername = document.getElementById('newUsername').value;
    const newTag = document.getElementById('newTag').value;
    if (!newTag.startsWith('#')) {
        const numericTag = newTag.replace(/[^0-9]/g, '');
        document.getElementById('usernameDisplay').textContent = `${newUsername}#${numericTag.padStart(4, '0').substring(0,4)}`;
    } else {
         document.getElementById('usernameDisplay').textContent = `${newUsername}${newTag}`;
    }

    console.log('Username saved:', newUsername + '#' + newTag);
    cancelEdit('username');
}

function saveEmail() {
    const newEmail = document.getElementById('newEmail').value;
    const atIndex = newEmail.indexOf('@');
    if (atIndex > 2) {
        const maskedEmail = newEmail.substring(0, 2) + '******' + newEmail.substring(atIndex);
        document.getElementById('emailDisplay').textContent = maskedEmail;
    } else {
        document.getElementById('emailDisplay').textContent = "em******@example.com";
    }

    console.log('Email saved:', newEmail);
    cancelEdit('email');
}

document.addEventListener('DOMContentLoaded', () => {
    console.log('Profile page loaded');
});
