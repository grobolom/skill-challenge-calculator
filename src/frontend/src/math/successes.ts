
// just used by the factorial function for fast calculations
const f: number[] = [];
function factorial(n: number): number {
  if (n == 0 || n == 1)
    return 1;
  if (f[n] > 0)
    return f[n];
  return f[n] = factorial(n-1) * n;
}

// This is a complicated function that's derived from the binomial
// distribution. It is used to calculate the odds of succeeding on a specific
// trial, given a total number of successes.
function calculateSingleProbability(baseProbability: number, successes: number, trials: number): number {
  const firstTerm = factorial(trials - 1) / (factorial(successes - 1) * factorial(trials - successes))
  const secondTerm = Math.pow(baseProbability, successes)
  const thirdTerm = Math.pow((1 - baseProbability), trials - successes)

  return firstTerm * secondTerm * thirdTerm
}

// Since the binomial probability calculation above only tackles specific trial
// instances (e.g. the probability of having 6 successes on the 8th trial), this
// function is needed to sum up the values for the whole range of failures and
// successes (e.g. the probability of getting the 6th success on the 6th, 7th,
// or 8th trial, for 6 successes and 3 failures as inputs)
function calculateProbability(baseProbability: number, successes: number, failures: number) {
  let sum = 0;

  for (let i = 0; i < failures; i++) {
    sum += calculateSingleProbability(baseProbability, successes, successes + i);
  }

  return sum
}

// Handles calculating probabilities for successes. Output is something like:
// [ { successes: 3, probability: 0.5 }, ... ]
// Currently uses a static list of successes to calculate for. The probability is
// calculated using the negative binomial distribution
export default function getSuccesses(
  skillBonus: number,
  checkDC: number,
  failures: number
): { successes: number; probability: number }[] {

  // the base probability for success is based on the character's skill bonus
  // compared to the check DC. A bonus of 5 and a DC of 16 will give a 50% chance
  // of success (0-10 fails, 11-20 succeeds).
  const baseProbability = (21 - (checkDC - skillBonus)) / 20;

  // just defining a constant here for convenience, but this could be expanded
  const successesRange = [3, 4, 5, 6, 7, 8, 9, 10, 11]

  return successesRange.map((value): { successes: number; probability: number } => {
    return {
      'successes': value,
      'probability': calculateProbability(baseProbability, value, failures)
    }
  });
}
