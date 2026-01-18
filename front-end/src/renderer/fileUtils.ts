/**
 * Small utility functions used by renderer components.
 * Keep these pure-ish helpers that don't depend on React state.
 */

export async function selectFile(): Promise<string | null> {
  try {
    // Call the exposed Electron API from preload
    const filePath = await (window as any).electronAPI.openFileDialog();

    if (filePath && typeof filePath === 'string') {
      return filePath;
    }

    return null;
  } catch (error) {
    console.error('Error opening file dialog:', error);
    return null;
  }
}
