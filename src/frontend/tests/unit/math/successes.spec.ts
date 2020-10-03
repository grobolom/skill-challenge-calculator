import { expect } from 'chai'
import getSuccesses from '@/math/successes'

describe('getSuccesses', () => {
  it('calculates a table of successes', () => {
    const skillBonus = 4
    const checkDC = 15
    const failures = 3

    const result = getSuccesses(skillBonus, checkDC, failures)

    expect(result[0].successes).to.eq(3)
    expect(result[0].probability).to.eq(0.500)

    expect(result[1].successes).to.eq(4)
    expect(result[1].probability).to.eq(0.34375)

    expect(result[2].successes).to.eq(5)
    expect(result[2].probability).to.eq(0.2265625)
  })
})