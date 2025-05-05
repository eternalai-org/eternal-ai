export const truncateAddress = (address: string) => {
  return address.slice(0, 6) + '...' + address.slice(-4);
};

export const formatNumber = (number: number) => {
  return number.toLocaleString('en-US', {
    maximumFractionDigits: 2,
  });
};
