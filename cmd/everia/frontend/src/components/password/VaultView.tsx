import React from 'react';

function VaultView({ entries }: { entries: any[] }) {
  return (
    <div className="p-4 max-w-2xl mx-auto">
      <h2 className="text-2xl font-semibold mb-4">🔐 Vault Entries</h2>
      <ul className="space-y-2">
        {entries.map((e, i) => (
          <li key={i} className="border p-2 rounded bg-gray-100">
            <strong>{e.site}</strong> <br />
            Username: {e.username} <br />
            Password: {e.password}
          </li>
        ))}
      </ul>
    </div>
  );
}

export default VaultView;
