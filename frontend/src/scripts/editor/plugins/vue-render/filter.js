export default function filter(str) {
  const replace = function(substr) {
    return substr.toLowerCase().replaceAll(/ /g, "-");
  };

  return Array.isArray(str) ? str.map(replace) : replace(str);
}
