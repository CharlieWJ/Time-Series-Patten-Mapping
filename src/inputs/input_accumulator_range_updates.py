accumulator_range_updates = {
    'C' : {
        'out_a' : {
            'a0' : ['default_g_f'],
            'a1' : ['default_g_f']
        },
        'found' : {
            'a0' : ['phi_f', '(', 'phi_f', '(', 'D', ',', 'delta_i_f', ')', ',', 'delta_i_f_prime', ') // C, found a0' ],
            'a1' : ['phi_f', '(', 'D', ',', 'delta_i_f', ') // C, found a1']
        },
        'in' : {
            'a0' : ['phi_f', '(', 'H', ',', 'delta_i_f_prime', ') // C, in a0'],
            'a1' : ['phi_f', '(', 'C', ',', 'phi_f', '(', 'D', ',', 'delta_i_f', ')', ') // C, in a1']
        }
    },
    'D' : {
        'out_r' : {
            'a0' : ['neutral_f'],
            'a1' : ['neutral_f']
        },
        'out_a' : {
            'a0' : ['neutral_f'],
            'a1' : ['neutral_f']
        },
        'maybe_b' : {
            'a0' : ['phi_f', '(', 'D', ',', 'delta_i_f', ')'],
            'a1' : ['phi_f', '(', 'D', ',', 'delta_i_f', ')']#dont alter
        },
        'maybe_a' : {
            'a0' : ['phi_f', '(', 'D', ',', 'delta_i_f_prime', ')'],
            'a1' : ['phi_f', '(', 'D', ',', 'delta_i_f', ')']
        },
        'found' : {
            'a0' : ['neutral_f'],
            'a1' : ['neutral_f']
        },
        'in' : {
            'a0' : ['neutral_f'],
            'a1' : ['neutral_f']
        },
        'found_e' : {
            'a0' : ['neutral_f'],
            'a1' : ['neutral_f']
        }
    },
    # NOTE: Added H
    'H' : {
        #'out_a' : {
        #    'a0' : ['default_g_f'],
        #    'a1' : ['default_g_f']
        #}
    },
    'R' : {
        'out_a' : {
            'a0' : ['g', '(', 'R', ',', 'C', ')'],
            'a1' : ['g', '(', 'R', ',', 'C', ')']
        },
        'found_e' : {
            'a0' : ['g', '(', 'R', ',', 'phi_f', '(', 'phi_f', '(', 'D', ',', 'delta_i_f', ')', ',', 'delta_i_f_prime', ')', ')'],
            'a1' : ['g', '(', 'R', ',', 'phi_f', '(', 'D', ',', 'delta_i_f', ')', ')']
        }
    }
}