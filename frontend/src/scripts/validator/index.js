export const validations = {
  domains:
    "((?=[a-z0-9-]{1,63}\\.)(xn--)?[a-z0-9]+(-[a-z0-9]+)*\\.)+[a-z]{2,63}",
  ips:
    "(?:(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?).){3}(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)",
  ports:
    "([0-9]{1,4}|[1-5][0-9]{4}|6[0-4][0-9]{3}|65[0-4][0-9]{2}|655[0-2][0-9]|6553[0-5])",
};

// extend validations
validations.withports = `(${validations.domains}|${validations.ips})(:${validations.ports})?`;
validations.baseurls = `[a-z]{2,63}:\/\/${validations.withports}`;
validations.urls = `${validations.baseurls}(\/.*)*`;

export default function validate(types, content) {
  if (!content || !content.length) return true;
  const checks = types.split(",");
  for (const check of checks) {
    const type = check.trim().toLowerCase();
    if (!validations[type] || content.match(`^${validations[type]}$`) !== null)
      return true;
  }
  return false;
}
