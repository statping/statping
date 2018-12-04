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

context('Settings Forms', () => {

    beforeEach(function() {
        cy.visit('http://localhost:8080/dashboard')
        cy.get('input[name="username"]').type('admin')
        cy.get('input[name="password"]').type('admin')
        cy.get('form').submit()
    })

    it('should edit main settings', () => {
        cy.visit('http://localhost:8080/settings')
        cy.get('input[name="project"]').clear().type('Project Updated')
        cy.get('input[name="description"]').clear().type('This is an awesome page')
        cy.get('input[name="domain"]').clear().type('http://0.0.0.0:8080')
        cy.get('textarea[name="footer"]').clear().type('This is a custom footer')
        cy.get('#v-pills-home > form').submit()
        cy.title().should('eq', 'Statping | Settings')
        cy.get('input[name="project"]').should('have.value', 'Project Updated')
        cy.get('input[name="description"]').should('have.value', 'This is an awesome page')
        cy.get('input[name="domain"]').should('have.value', 'http://0.0.0.0:8080')
        cy.get('.footer').should('contain', 'This is a custom footer')
    })

    // it('should check index page for changes', () => {
    //     cy.visit('http://localhost:8080/')
    //     cy.title().should('eq', 'Project Updated Status')
    //     cy.get('.header-title').should('contain', 'Project Updated')
    //     cy.get('.header-desc').should('contain', 'This is an awesome page')
    // })

});
