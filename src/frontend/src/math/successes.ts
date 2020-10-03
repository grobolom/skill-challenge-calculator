
// just used by the factorial function for fast calculations
const f: number[] = [];
function factorial(n: number): number {
  if (n == 0 || n == 1)
    return 1;
  if (f[n] > 0)
    return f[n];
  return f[n] = factorial(n-1) * n;
}

function calculateSingleProbability(baseProbability: number, successes: number, trials: number): number {
  const firstTerm = factorial(trials - 1) / (factorial(successes - 1) * factorial(trials - successes))
  const secondTerm = Math.pow(baseProbability, successes)
  const thirdTerm = Math.pow((1 - baseProbability), trials - successes)

  return firstTerm * secondTerm * thirdTerm
}

function calculateProbability(baseProbability: number, successes: number, failures: number) {
  let sum = 0;

  for (let i = 0; i < failures; i++) {
    sum += calculateSingleProbability(baseProbability, successes, successes + i);
  }

  return sum
}

export default function getSuccesses(skillBonus: number, checkDC: number, failures: number): { successes: number; probability: number }[] {
  const baseProbability = (21 - (checkDC - skillBonus)) / 20;

  // just defining this here for convenience
  const successesRange = [3, 4, 5, 6, 7, 8, 9, 10, 11]

  return successesRange.map((value): { successes: number; probability: number } => {
    return {
      'successes': value,
      'probability': calculateProbability(baseProbability, value, failures)
    }
  });
}
