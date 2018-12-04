/*
 * Statping
 * Copyright (C) 2018.  Hunter Long and the project contributors
 * Written by Hunter Long <info@socialeck.com> and the project contributors
 *
 * https://github.com/hunterlong/statping
 *
 * The licenses for most software and other practical works are designed
 * to take away your freedom to share and change the works.  By contrast,
 * the GNU General Public License is intended to guarantee your freedom to
 * share and change all versions of a program--to make sure it remains free
 * software for all its users.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

context('Setup Process', () => {

    // it('should go to setup Statping with Postgres', () => {
    //     cy.visit('http://localhost:8080')
    //     cy.get('select[name=db_connection]').select('postgres')
    //     cy.get('input[name="db_host"]').clear().type(Cypress.env('DB_HOST'))
    //     cy.get('input[name="db_port"]').clear().type('5432')
    //     cy.get('input[name="db_user"]').clear().type(Cypress.env('DB_USER'))
    //     if (Cypress.env('TRAVIS')==="yes") {
    //         cy.get('input[name="db_password"]').clear()
    //     } else {
    //         cy.get('input[name="db_password"]').clear().type(Cypress.env('DB_PASS'))
    //     }
    //     cy.get('input[name="db_database"]').clear().type(Cypress.env('DB_DATABASE'))
    //     cy.get('input[name="project"]').clear().type('Demo Tester')
    //     cy.get('input[name="description"]').clear().type('This is a test from Crypress!')
    //     cy.get('input[name="domain"]').clear().type('http://localhost:8080')
    //     cy.get('input[name="username"]').clear().type('admin')
    //     cy.get('input[name="email"]').clear().type('info@domain.com')
    //     cy.get('input[name="password"]').clear().type('admin')
    //     cy.scrollTo('bottom')
    //     cy.get('#setup_button').click().wait(10000)
    //     cy.get('.header-title').should('contain', 'Demo Tester')
    //     cy.get('.header-desc').should('contain', 'This is a test from Crypress!')
    //     cy.scrollTo('bottom')
    //     cy.get('.service_li').should('have.length', 5)
    //     cy.get('.card').should('have.length', 5)
    // })

    it('should go to setup Statping with SQLite', () => {
        cy.visit('http://localhost:8080')
        cy.get('select[name=db_connection]').select('sqlite')
        cy.get('input[name="project"]').clear().type('Demo Tester')
        cy.get('input[name="description"]').clear().type('This is a test from Crypress!')
        cy.get('input[name="domain"]').clear().type('http://localhost:8080')
        cy.get('input[name="username"]').clear().type('admin')
        cy.get('input[name="email"]').clear().type('info@domain.com')
        cy.get('input[name="password"]').clear().type('admin')
        cy.scrollTo('bottom')
        cy.get('#setup_button').click()
        cy.get('.header-title').should('contain', 'Demo Tester')
        cy.get('.header-desc').should('contain', 'This is a test from Crypress!')
        cy.scrollTo('bottom')
        cy.get('.service_li').should('have.length', 5)
        cy.get('.card').should('have.length', 5)
    })

})
