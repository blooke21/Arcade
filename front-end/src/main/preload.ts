const { contextBridge, ipcRenderer } = require('electron');

console.log('Preload script is running!');

// Expose protected methods that allow the renderer process to use
// Electron APIs without exposing the entire object
contextBridge.exposeInMainWorld('electronAPI', {
  openFileDialog: () => ipcRenderer.invoke('dialog:openFile'),
});
