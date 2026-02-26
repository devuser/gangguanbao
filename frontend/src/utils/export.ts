export function downloadBlob(blob: Blob, filename: string) {
  const url = URL.createObjectURL(blob);
  const a = document.createElement('a');
  a.href = url;
  a.download = filename;
  a.click();
  URL.revokeObjectURL(url);
}

export function getExportFilename(contentDisposition: string): string {
  const match = contentDisposition?.match(/filename[^;=\n]*=((['"]).*?\2|[^;\n]*)/);
  if (match) {
    return decodeURIComponent(match[1].replace(/['"]/g, ''));
  }
  return `export_${Date.now()}.xlsx`;
}
