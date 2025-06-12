document.addEventListener('DOMContentLoaded', () => {
    const createMyOwnButton = document.getElementById('createMyOwn');
    const serverCreationForm = document.getElementById('serverCreationForm');
    const templatesSection = document.querySelector('.templates-section'); // Assuming this and initial h1/p are part of the initial view
    const initialContent = [document.querySelector('.add-server-card > h1'), document.querySelector('.add-server-card > p'), createMyOwnButton, templatesSection];

    const backToTemplatesButton = document.getElementById('backToTemplatesButton');
    const createServerButton = document.getElementById('createServerButton');
    const serverNameInput = document.getElementById('serverNameInput');
    const iconUploader = document.getElementById('iconUploader');
    const serverIconInput = document.getElementById('serverIconInput');
    const serverIconPreview = document.getElementById('serverIconPreview');

    if (createMyOwnButton) {
        createMyOwnButton.addEventListener('click', () => {
            initialContent.forEach(el => el.style.display = 'none');
            if (serverCreationForm) serverCreationForm.style.display = 'block';
        });
    }

    if (backToTemplatesButton) {
        backToTemplatesButton.addEventListener('click', () => {
            if (serverCreationForm) serverCreationForm.style.display = 'none';
            initialContent.forEach(el => el.style.display = el.tagName === 'DIV' || el.tagName === 'H1' || el.tagName === 'P' ? 'block' : 'flex'); // Adjust display based on original
             // Specifically for createMyOwnButton if it's flex
            if(createMyOwnButton) createMyOwnButton.style.display = 'flex';
        });
    }

    if (createServerButton && serverNameInput) {
        createServerButton.addEventListener('click', () => {
            const serverName = serverNameInput.value.trim() || "Nouveau Serveur";
            // For now, redirect to a generic server page with the name as a query parameter
            // In a real app, this would involve an API call to create the server.
            window.location.href = `/views/server-view.html?name=${encodeURIComponent(serverName)}`;
        });
    }
    
    if (iconUploader && serverIconInput) {
        iconUploader.addEventListener('click', () => serverIconInput.click());
        serverIconInput.addEventListener('change', (event) => {
            const file = event.target.files[0];
            if (file && serverIconPreview) {
                const reader = new FileReader();
                reader.onload = (e) => {
                    serverIconPreview.src = e.target.result;
                }
                reader.readAsDataURL(file);
            }
        });
    }
});
