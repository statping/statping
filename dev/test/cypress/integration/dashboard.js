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

context('Dashboard Tests', () => {

    beforeEach(function() {
        cy.visit('http://localhost:8080/dashboard')
        cy.get('input[name="username"]').type('admin')
        cy.get('input[name="password"]').type('admin')
        cy.get('form').submit()
    })

    it('should view logs', () => {
        cy.visit('http://localhost:8080/settings')
        cy.get(':nth-child(5) > .nav-link').click()
        cy.wait(10000)
        cy.get('#live_logs').should('contain', 'Service')
    })

    it('should view help', () => {
        cy.visit('http://localhost:8080/settings')
        cy.get(':nth-child(6) > .nav-link').click()
        cy.title().should('eq', 'Statping | Help')
        cy.get('.col-12 > :nth-child(1)').should('contain', 'Statping')
    })

});
